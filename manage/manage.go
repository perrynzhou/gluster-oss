package manage

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	fs_api "glusterfs-storage-gateway/fs-api"
	"glusterfs-storage-gateway/meta"
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
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		buketMeta := scanner.Text()
		buketMeta = buketMeta[0 : len(buketMeta)-1]
		bucket := &Bucket{}
		if err := json.Unmarshal([]byte(buketMeta), bucket); err != nil {
			return nil, err
		}
		bucket.Load(api)
		bucketManage.BucketCache[bucket.Meta.Name] = bucket
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
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
		case <-bucketManage.doneCh:
			return
		}
	}
}
func (bucketManage *BucketManage) handleCreateBucketRequest(request *BucketRequest) error {
	var err error
	var bucket *Bucket
	response := &bucket.BucketResponse{}
	defer bucketManage.handleError(request, response, err)
	if bucket, err = bucketManage.checkBucket(request.Info.Name); err != nil {
		log.Infoln("handleCreateBucketRequest fetch request:", request)
		bucketDirName := fmt.Sprintf("%s-%s", request.Info.Name, uuid.New().String())
		if bucket = NewBucket(bucketManage.api, request.Info.LimitSize, request.Info.LimitCount, request.Info.Name, bucketDirName); bucket == nil {
			response.Err = errors.New("create bucket faild")
		}
		bucketManage.notifyCh <- bucket
		return nil
	}
	return err
}
func (bucketManage *BucketManage) handleError(request *BucketRequest, response *BucketResponse, err error) {
	response.Reply = request.Info
	response.Err = err
	request.Done <- response
}
func (bucketManage *BucketManage) checkBucket(bucketName string) (*Bucket, error) {
	var bucket *Bucket
	var err error
	var ok bool
	// check bucket is exists
	if bucket, ok = bucketManage.BucketCache[bucketName]; !ok {
		err = errors.New(fmt.Sprintf("%s not exists", bucketName))
		return nil, err
	}
	// check bucket is invalid
	if err = bucket.IsPermit(); err != nil {
		return nil, err
	}
	return bucket, nil
}
func (bucketManage *BucketManage) handleUpdateBucketRequest(request *BucketRequest) error {
	var bucket *Bucket
	var err error
	var isChange bool
	response := &bucket.BucketResponse{}
	defer bucketManage.handleError(request, response, err)
	if bucket, err = bucketManage.checkBucket(request.Info.Name); err != nil {
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
		StoreBucket(bucketManage.api, bucketManage.bucketMetaFile, bucket)
		bucketManage.notifyCh <- bucket
	}
	return nil
}
func (bucketManage *BucketManage) handleDeleteBucketRequest(request *BucketRequest) error {
	var err error
	var bucketInfo *meta.BucketInfo
	var bucket *Bucket
	bucketInfo, err = bucketManage.fetchBucketInfo(request.Info.Name)
	response := &bucket.BucketResponse{}
	defer bucketManage.handleError(request, response, err)
	bucket, err = bucketManage.checkBucket(request.Info.Name)
	if err != nil {
		return err
	}
	bucketInfo.Status = bucket.BucketInActiveStatus
	StoreBucket(bucketManage.api, bucketManage.bucketMetaFile, bucket)
	go bucketManage.delBucketInfoAndBucketData(request, bucketInfo)
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
