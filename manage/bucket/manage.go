package bucket

/*
  gluster volume : /
              /bucket.info
              /bucket1
              /bucekt2
*/
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
	manage := &BucketManage{
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
	if err = manage.api.Stat(bucketInfoFilePath); err != nil {
		bucketInfoFile, err = manage.api.Creat(bucketInfoFilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	} else {
		bucketInfoFile, err = manage.api.Open(bucketInfoFilePath, os.O_RDWR|os.O_APPEND)
	}
	if err != nil {
		return nil, err
	}
	manage.bucketInfoFile = bucketInfoFile
	return manage, nil
}

func (manage *BucketManage) refreshCache() {
	log.Infoln("run BucketService refreshCache")
	defer manage.wg.Done()
	for {
		select {
		case bucketInfo := <-manage.notifyCh:
			manage.BucketInfoCache[bucketInfo.Name] = bucketInfo
		case <-manage.doneCh:
			return
		}
	}
}
func (manage *BucketManage) handleCreateBucketRequest(request *BucketInfoRequest) error {
	response := &BucketInfoResponse{
		Reply: request.Info,
	}

	defer func(request *BucketInfoRequest, response *BucketInfoResponse) {
		if response.Err != nil {
			log.Errorln(response.Err)
		}
		request.Done <- response
	}(request, response)
	if manage.checkBucketExist(request.Info.Name) {
		response.Err = errors.New(fmt.Sprintf("%s already exists", request.Info.Name))
		return response.Err
	}
	log.Infoln("handleCreateBucketRequest fetch request:", request)
	var b []byte
	if err := manage.handleBucketDir(request.Info.Name, request.Info.RealDirName, createBucketDirType); err != nil {
		response.Err = err
	} else {
		if b, err = manage.storeBucketInfo(request.Info, request.RequestType); err != nil {
			manage.handleBucketDir(request.Info.Name, request.Info.RealDirName, deleteBucketDirType)
			request.Info.Status = DeleteBucketType
			manage.conn.Del(context.Background(), request.Info.Name)
			response.Err = err
		} else {
			response.Err = nil
			log.Infoln("handleCreateBucketRequest resp:::", response)
			manage.conn.Set(context.Background(), request.Info.Name, string(b), -1)
			manage.notifyCh <- request.Info
		}
	}

	return response.Err
}
func (manage *BucketManage) handleUpdateBucketRequest(request *BucketInfoRequest) error {

	bucketInfo, err := manage.fetchBucketInfo(request.Info.Name)
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
	if b, err = manage.storeBucketInfo(bucketInfo, request.RequestType); err != nil {
		response.Err = err
		return err
	}
	manage.conn.Set(context.Background(), request.Info.Name, string(b), -1)
	manage.persistenceBucketInfoToDisk(request.Info.Name, b)
	manage.notifyCh <- request.Info
	return nil
}
func (manage *BucketManage) handleDeleteBucketRequest(request *BucketInfoRequest) error {
	bucketInfo, err := manage.fetchBucketInfo(request.Info.Name)
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
	manage.storeBucketInfo(bucketInfo, DeleteBucketType)
	go manage.delBucketInfoAndBucketData(request, bucketInfo)
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
	manage.wg.Done()
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
			}
		case <-manage.doneCh:
			return
		}
	}
}
func (manage *BucketManage) Stop() {
	for i := 0; i < manage.goFuncCount; i++ {
		manage.doneCh <- struct{}{}
	}
}
