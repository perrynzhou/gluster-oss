package bucket

import (
	"errors"
	"fmt"
	fs_api "glusterfs-storage-gateway/fs-api"
	"glusterfs-storage-gateway/meta"
	"os"
	"sync"

	"golang.org/x/net/context"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

const (
	bucketInfoFilePath = "/bucket.idx"
)

type BucketManage struct {
	api             *fs_api.FsApi
	conn            *redis.Conn
	ReqCh           chan *BucketInfoRequest
	doneCh          chan struct{}
	wg              *sync.WaitGroup
	BucketInfoCache map[string]*meta.BucketInfo
	notifyCh        chan *meta.BucketInfo
	goFuncCount     int
	bucketInfoFile  *fs_api.FsFd
}

func NewBucketManage(api *fs_api.FsApi, conn *redis.Conn, bucketRequestCh chan *BucketInfoRequest, wg *sync.WaitGroup) (*BucketManage, error) {
	bucketManage := &BucketManage{
		api:             api,
		conn:            conn,
		ReqCh:           bucketRequestCh,
		wg:              wg,
		notifyCh:        make(chan *meta.BucketInfo),
		doneCh:          make(chan struct{}),
		BucketInfoCache: make(map[string]*meta.BucketInfo, 0),
		goFuncCount:     0,
	}
	var bucketInfoFile *fs_api.FsFd
	var err error
	if err = bucketManage.api.Stat(bucketInfoFilePath); err != nil {
		bucketInfoFile, err = bucketManage.api.Creat(bucketInfoFilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	} else {
		bucketInfoFile, err = bucketManage.api.Open(bucketInfoFilePath, os.O_RDWR|os.O_APPEND)
	}
	if err != nil {
		return nil, err
	}
	bucketManage.bucketInfoFile = bucketInfoFile
	return bucketManage, nil
}

func (bucketManage *BucketManage) refreshCache() {
	log.Infoln("run BucketService refreshCache")
	defer bucketManage.wg.Done()
	for {
		select {
		case bucketInfo := <-bucketManage.notifyCh:
			bucketManage.BucketInfoCache[bucketInfo.Name] = bucketInfo
		case <-bucketManage.doneCh:
			return
		}
	}
}
func (bucketManage *BucketManage) handleCreateBucketRequest(request *BucketInfoRequest) error {
	response := &BucketInfoResponse{
		Reply: request.Info,
	}

	defer func(request *BucketInfoRequest, response *BucketInfoResponse) {
		if response.Err != nil {
			log.Errorln(response.Err)
		}
		request.Done <- response
	}(request, response)
	if bucketManage.checkBucketExist(request.Info.Name) {
		response.Err = errors.New(fmt.Sprintf("%s already exists", request.Info.Name))
		return response.Err
	}
	log.Infoln("handleCreateBucketRequest fetch request:", request)
	var b []byte
	if err := bucketManage.handleBucketDir(request.Info.Name, request.Info.RealDirName, createBucketDirType); err != nil {
		response.Err = err
	} else {
		if b, err = bucketManage.storeBucketInfo(request.Info, request.RequestType); err != nil {
			bucketManage.handleBucketDir(request.Info.Name, request.Info.RealDirName, deleteBucketDirType)
			request.Info.Status = DeleteBucketType
			bucketManage.conn.Del(context.Background(), request.Info.Name)
			response.Err = err
		} else {
			response.Err = nil
			log.Infoln("handleCreateBucketRequest resp:::", response)
			bucketManage.conn.Set(context.Background(), request.Info.Name, string(b), -1)
			bucketManage.notifyCh <- request.Info
		}
	}

	return response.Err
}
func (bucketManage *BucketManage) handleUpdateBucketRequest(request *BucketInfoRequest) error {

	bucketInfo, err := bucketManage.fetchBucketInfo(request.Info.Name)
	response := &BucketInfoResponse{}
	defer func(request *BucketInfoRequest, response *BucketInfoResponse, bucketInfo *meta.BucketInfo) {
		request.Done <- response
		response.Reply = bucketInfo
	}(request, response, bucketInfo)
	if err != nil {
		response.Err = err
		return err
	}
	response.Reply = bucketInfo
	if bucketInfo.UsageInfo.CapacityLimitSize > request.Info.UsageInfo.CapacityLimitSize {
		response.Err = errors.New(fmt.Sprintf("invalid CapacityLimitSize(%d<=%d)", bucketInfo.UsageInfo.CapacityLimitSize, request.Info.UsageInfo.CapacityLimitSize))
	} else if bucketInfo.UsageInfo.ObjectsLimitCount > request.Info.UsageInfo.ObjectsLimitCount {
		response.Err = errors.New(fmt.Sprintf("invalid ObjectsLimitCount(%d<=%d)", bucketInfo.UsageInfo.ObjectsLimitCount, request.Info.UsageInfo.ObjectsLimitCount))
	}
	if response.Err != nil {
		return response.Err
	}
	bucketInfo.UsageInfo.ObjectsLimitCount = request.Info.UsageInfo.ObjectsLimitCount
	bucketInfo.UsageInfo.CapacityLimitSize = request.Info.UsageInfo.CapacityLimitSize
	var b []byte
	if b, err = bucketManage.storeBucketInfo(bucketInfo, request.RequestType); err != nil {
		response.Err = err
		return err
	}
	bucketManage.conn.Set(context.Background(), request.Info.Name, string(b), -1)
	bucketManage.persistenceBucketInfoToDisk(request.Info.Name, b)
	bucketManage.notifyCh <- request.Info
	return nil
}
func (bucketManage *BucketManage) handleDeleteBucketRequest(request *BucketInfoRequest) error {
	bucketInfo, err := bucketManage.fetchBucketInfo(request.Info.Name)
	response := &BucketInfoResponse{}
	defer func(request *BucketInfoRequest, response *BucketInfoResponse, bucketInfo *meta.BucketInfo) {
		request.Done <- response
		response.Reply = bucketInfo
	}(request, response, bucketInfo)
	if err != nil {
		response.Err = err
		return err
	}
	bucketInfo.Status = BucketInActiveStatus
	bucketManage.storeBucketInfo(bucketInfo, DeleteBucketType)
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
