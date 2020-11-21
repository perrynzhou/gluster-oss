package bucket

import (
	"errors"
	"fmt"
	fs_api "glusterfs-storage-gateway/fs-api"
	"glusterfs-storage-gateway/meta"
	"sync"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

type BucketManage struct {
	api             *fs_api.FsApi
	conn            *redis.Conn
	ReqCh           chan *BucketInfoRequest
	doneCh          chan struct{}
	wg              *sync.WaitGroup
	bucketInfoCache map[string]*meta.BucketInfo
	notifyCh        chan *meta.BucketInfo
	goFuncCount     int
}

func NewBucketManage(api *fs_api.FsApi, conn *redis.Conn,bucketRequestCh chan *BucketInfoRequest, wg *sync.WaitGroup) *BucketManage {
	return &BucketManage{
		api:         api,
		conn:        conn,
		ReqCh:      bucketRequestCh,
		wg:          wg,
		notifyCh:    make(chan *meta.BucketInfo),
		doneCh:      make(chan struct{}),
		bucketInfoCache: make(map[string]*meta.BucketInfo, 0),
		goFuncCount: 0,
	}
}
func (manage *BucketManage) refreshCache() {
	log.Infoln("run BucketService refreshCache")
	defer manage.wg.Done()
	for {
		select {
		case bucketInfo := <-manage.notifyCh:
			manage.bucketInfoCache[bucketInfo.Name] = bucketInfo
			break
		case <-manage.doneCh:
			return
		default:
			break
		}
	}
}
func (manage *BucketManage) handleCreateBucketRequest(request *BucketInfoRequest) error {
	response := &BucketInfoResponse{
		Reply:request.Info,
	}
	defer func(request *BucketInfoRequest,response *BucketInfoResponse) {
		request.Done <- response
	}(request,response)
	log.Infoln("handleCreateBucketRequest fetch request:",request)
	if !manage.checkBucketExist(request.Info.Name){
		if err := manage.handleBucketDir(request.Info.Name,request.Info.RealDirName, createBucketDirType); err != nil {
			response.Err = err
		} else {
			if _, err := manage.storeBucketInfo(request.Info); err != nil {
				manage.handleBucketDir(request.Info.Name, request.Info.RealDirName, deleteBucketDirType)
				response.Err = err
			} else {
				response.Err = nil
				log.Infoln("handleCreateBucketRequest resp:::",response)
				manage.notifyCh <- request.Info
			}
		}
	}
	if response.Err != nil{
		log.Errorln("handleCreateBucketRequest err:",response.Err)
	}

	return response.Err
}
func (manage *BucketManage) handleUpdateBucketRequest(request *BucketInfoRequest) error {
	bucketInfo :=request.Info
	response := &BucketInfoResponse{
		Reply:bucketInfo,
	}
	bucketInfo, err := manage.fetchBucketInfo(request.Info.Name)
	defer func(request *BucketInfoRequest,response *BucketInfoResponse) {
		request.Done <- response
	}(request,response)
	if err != nil {
		response.Err = err
		return err
	}
	if bucketInfo.UsageInfo.CapacityLimitSize <= request.Info.UsageInfo.CapacityLimitSize {
		response.Err = errors.New(fmt.Sprintf("invalid CapacityLimitSize(%d<=%d)", bucketInfo.UsageInfo.CapacityLimitSize, request.Info.UsageInfo.CapacityLimitSize))
	} else if bucketInfo.UsageInfo.ObjectsLimitCount <= request.Info.UsageInfo.ObjectsLimitCount {
		response.Err = errors.New(fmt.Sprintf("invalid ObjectsLimitCount(%d<=%d)", bucketInfo.UsageInfo.ObjectsLimitCount, request.Info.UsageInfo.ObjectsLimitCount))
	}
	if response.Err != nil {
		return response.Err
	}
	bucketInfo.UsageInfo.ObjectsLimitCount = request.Info.UsageInfo.ObjectsLimitCount
	bucketInfo.UsageInfo.CapacityLimitSize = request.Info.UsageInfo.CapacityLimitSize
	if _, err := manage.storeBucketInfo(bucketInfo); err != nil {
		response.Err = err
		return err
	}
	manage.notifyCh <- request.Info
	return nil
}
func (manage *BucketManage) handleDeleteBucketRequest(request *BucketInfoRequest) error {
	bucketInfo :=request.Info
	response := &BucketInfoResponse{
		Reply:bucketInfo,
	}
	bucketInfo, err := manage.fetchBucketInfo(request.Info.Name)
	if err != nil {
		response.Err = err
		request.Done <- response
		return err
	}
	go manage.delBucketInfoAndBucketData(request, bucketInfo.RealDirName)
	return nil
}
func (manage *BucketManage) Run() {
	manage.goFuncCount = 2
	manage.wg.Add(manage.goFuncCount)
	go manage.handleBucketRequest()
	go manage.refreshCache()

}
func (manage *BucketManage) handleBucketRequest() {
	log.Infoln("run BucketService handleBucketRequest")
	defer manage.wg.Done()
	for {
		select {
		case req := <-manage.ReqCh:
			log.Infoln("recive request:", req)
			switch req.RequestType {
			case CreateBucketType:
				manage.handleCreateBucketRequest(req)
				break
			case DeleteBucketType:
				manage.handleDeleteBucketRequest(req)
				break
			case UpdateBucketType:
				manage.handleUpdateBucketRequest(req)
				break
			default:
				break
			}
			break
		case <-manage.doneCh:
			return
		default:
			break
		}
	}
}
func (manage *BucketManage) Stop() {
	for i := 0; i < manage.goFuncCount; i++ {
		manage.doneCh <- struct{}{}
	}
}
