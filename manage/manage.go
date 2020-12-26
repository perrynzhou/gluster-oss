package manage

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	fs_api "glusterfs-storage-gateway/fs-api"
	"glusterfs-storage-gateway/meta"
	"io"
	"os"
	"sync"

	"github.com/google/uuid"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

const (
	bucketMetaFilePath = "/bucket.meta"
)

type BucketManage struct {
	api            *fs_api.FsApi
	conn           *redis.Conn
	ReqCh          chan *BucketRequest
	doneCh         chan struct{}
	wg             *sync.WaitGroup
	BucketCache    map[string]*Bucket
	notifyCh       chan *Bucket
	goFuncCount    int
	bucketMetaFile *fs_api.FsFd
}

func initBucketManage(api *fs_api.FsApi, bucketRequestCh chan *BucketRequest, wg *sync.WaitGroup) *BucketManage {
	return &BucketManage{
		api:         api,
		ReqCh:       bucketRequestCh,
		wg:          wg,
		notifyCh:    make(chan *Bucket),
		doneCh:      make(chan struct{}),
		BucketCache: make(map[string]*Bucket, 0),
		goFuncCount: 0,
	}
}
func InitBucketManage(api *fs_api.FsApi, bucketMetaFilePath string, bucketRequestCh chan *BucketRequest, wg *sync.WaitGroup) (*BucketManage, error) {
	bucketMetaFile, err := api.Open(bucketMetaFilePath, os.O_RDWR|os.O_APPEND)
	if err != nil {
		return nil, err
	}
	bucketManage := initBucketManage(api, bucketRequestCh, wg)
	bucketManage.bucketMetaFile = bucketMetaFile
	data := make([]byte, 1024*1024*256)
	api.Read(bucketMetaFile, data)
	buf := bytes.NewBuffer(data)
	reader := bufio.NewReader(buf)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		bucket := &Bucket{}
		if err := json.Unmarshal([]byte(data), bucket); err != nil {
			return nil, err
		}
		bucket.Load(api)
		bucketManage.BucketCache[bucket.Meta.Name] = bucket
		// Process the line here.
		log.Printf(" > Read %d characters\n", len(line))
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		log.Printf(" > Failed with error: %v\n", err)
		return nil, err
	}
	return bucketManage, nil
}
func NewBucketManage(api *fs_api.FsApi, bucketRequestCh chan *BucketRequest, wg *sync.WaitGroup) (*BucketManage, error) {
	var err error
	bucketManage := initBucketManage(api, bucketRequestCh, wg)
	if err = bucketManage.api.Stat(bucketMetaFilePath); err != nil {
		var bucketMetaFile *fs_api.FsFd
		bucketMetaFile, err = bucketManage.api.Creat(bucketMetaFilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
		bucketManage.bucketMetaFile = bucketMetaFile
	} else {
		return InitBucketManage(api, bucketMetaFilePath, bucketRequestCh, wg)
	}
	return bucketManage, nil
}

func (bucketManage *BucketManage) refreshCache() {
	log.Infoln("run BucketService refreshCache")
	defer bucketManage.wg.Done()
	for {
		select {
		case bucket := <-bucketManage.notifyCh:
			bucketManage.BucketCache[bucket.Meta.Name] = bucket
			bucket.StoreMeta(bucketManage.api, bucketManage.bucketMetaFile)
			break
		case <-bucketManage.doneCh:
			return
		}
	}
}
func (bucketManage *BucketManage) handleCreateBucketRequest(request *BucketRequest) error {
	var err error
	var bucket *Bucket
	var ok bool
	response := &BucketResponse{
		Reply: request.Info,
	}
	defer func(request *BucketRequest) {
		request.Done <- response
	}(request)
	if bucket, ok = bucketManage.BucketCache[request.Info.Name]; ok {
		err = errors.New(fmt.Sprintf("bucket %s  exists", request.Info.Name))
		response.Err = err
		return err
	}
	log.Infoln("handleCreateBucketRequest fetch request:", request)
	bucketDirName := fmt.Sprintf("%s-%s", request.Info.Name, uuid.New().String())
	if bucket, err = NewBucket(bucketManage.api, request.Info.LimitSize, request.Info.LimitCount, request.Info.Name, bucketDirName); err != nil {
		err = errors.New("create bucket faild")
		return err
	}
	bucketManage.notifyCh <- bucket
	return nil

}
func (bucketManage *BucketManage) handleError(request *BucketRequest, response *BucketResponse, err error) {
	response.Reply = request.Info
	response.Err = err
	if err != nil {
		log.Errorln("happen err:", err)
	}
	request.Done <- response
}
func (bucketManage *BucketManage) handleUpdateBucketRequest(request *BucketRequest) error {
	var bucket *Bucket
	var err error
	var isChange bool
	var ok bool
	response := &BucketResponse{
		Reply: request.Info,
		Err:   nil,
	}
	defer func(request *BucketRequest) {
		request.Done <- response
	}(request)
	if bucket, ok = bucketManage.BucketCache[request.Info.Name]; !ok {
		err = errors.New(fmt.Sprintf("bucket %s not exists", request.Info.Name))
		response.Err = err
		return err
	}

	if err = bucket.checkStatus(); err != nil {
		err = errors.New(fmt.Sprintf("bucket %s is  inactive", request.Info.Name))
		response.Err = err
		return err
	}
	if err = bucket.CheckLimit(); err != nil {
		err = errors.New(fmt.Sprintf("bucket %s over limits", request.Info.Name))
		response.Err = err
		return err
	}
	if bucket.Meta.LimitCount < request.Info.LimitCount {
		bucket.Meta.LimitCount = request.Info.LimitCount
		isChange = true
	}
	if bucket.Meta.LimitSize < request.Info.LimitSize {
		bucket.Meta.LimitSize = request.Info.LimitSize
		isChange = true
	}
	if isChange {
		bucketManage.notifyCh <- bucket
	}

	return nil
}
func (bucketManage *BucketManage) handleDeleteBucketRequest(request *BucketRequest) error {
	var err error
	var bucket *Bucket
	var ok bool
	response := &BucketResponse{
		Reply: request.Info,
	}
	defer func(request *BucketRequest) {
		request.Done <- response
	}(request)
	if bucket, ok = bucketManage.BucketCache[request.Info.Name]; !ok {
		err = errors.New(fmt.Sprintf("bucket %s not exists", request.Info.Name))
		response.Err = err
		return err
	}
	if err = bucket.checkStatus(); err != nil {
		response.Err = err
		return err
	}
	if err = bucket.CheckLimit(); err != nil {
		response.Err = err
		return err
	}
	bucket.Meta.Status = meta.InactiveBucket
	bucketManage.notifyCh <- bucket
	return nil
}
func (bucketManage *BucketManage) Run() {
	bucketManage.goFuncCount = 2
	bucketManage.wg.Add(bucketManage.goFuncCount)
	go bucketManage.handleBucketRequest()
	go bucketManage.refreshCache()

}
func (bucketManage *BucketManage) handleBucketRequest() {
	log.Infoln("run BucketService handleBucketRequest")
	bucketManage.wg.Done()
	for {
		select {
		case req := <-bucketManage.ReqCh:
			log.Infoln("recive request:", req)
			switch req.RequestType {
			case CreateBucketType:
				bucketManage.handleCreateBucketRequest(req)
				break
			case DeleteBucketType:
				bucketManage.handleDeleteBucketRequest(req)
				break
			case UpdateBucketType:
				bucketManage.handleUpdateBucketRequest(req)
				break
			}
		case <-bucketManage.doneCh:
			return
		}
	}
}
func (bucketManage *BucketManage) Stop() {
	for i := 0; i < bucketManage.goFuncCount; i++ {
		bucketManage.doneCh <- struct{}{}
	}
}
