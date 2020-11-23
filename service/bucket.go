package service

import (
	fs_api "glusterfs-storage-gateway/fs-api"
	"glusterfs-storage-gateway/manage/bucket"
	"glusterfs-storage-gateway/meta"
	"glusterfs-storage-gateway/protocol/pb"
	"glusterfs-storage-gateway/utils"
	"sync"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type BucketService struct {
	fsApi           *fs_api.FsApi
	ServiceName     string
	bucketRequestCh chan *bucket.BucketInfoRequest
	bucketMange     *bucket.BucketManage
	wg              *sync.WaitGroup
}

func NewBucketSerivce(api *fs_api.FsApi, serviceName string, wg *sync.WaitGroup) *BucketService {
	var err error
	bucketService := &BucketService{
		fsApi:           api,
		ServiceName:     serviceName,
		bucketRequestCh: make(chan *bucket.BucketInfoRequest),
		wg:              wg,
	}
	redisCon := utils.RedisClient.Conn(context.Background())
	bucketService.bucketMange, err = bucket.NewBucketManage(api, redisCon, bucketService.bucketRequestCh, wg)
	if err != nil {
		return nil
	}
	log.Info("init BucketService success")
	return bucketService
}

func (s *BucketService) Run() {
	s.bucketMange.Run()
}
func (s *BucketService) Stop() {
	s.bucketMange.Stop()
}
func (s *BucketService) CreateBucket(ctx context.Context, createBucketRequest *pb.CreateBucketRequest) (*pb.CreateBucketResponse, error) {
	req := bucket.NewCreateBucketInfoRequest(createBucketRequest)
	log.Infoln("get CreateBucket request:", createBucketRequest)
	s.bucketRequestCh <- req
	resp := <-req.Done

	bucketInfo := resp.Reply.(*meta.BucketInfo)
	log.Infoln("finish CreateBucket request:", bucketInfo, ",err:", resp.Err)
	createBucketResponse := &pb.CreateBucketResponse{
		Name: bucketInfo.Name,
		//request storage capacity
		Capacity: bucketInfo.UsageInfo.CapacityLimitSize,
		//obejcts limits
		ObjectsLimit: bucketInfo.UsageInfo.ObjectsLimitCount,
		BucketDir:    bucketInfo.RealDirName,
		Message:      "success",
	}
	if resp.Err != nil {
		createBucketResponse.Message = resp.Err.Error()
	}
	return createBucketResponse, resp.Err
}

func (s *BucketService) DeleteBucket(ctx context.Context, deleteBucketRequest *pb.DeleteBucketRequest) (*pb.DeleteBucketResponse, error) {
	req := bucket.NewDeleteBucketInfoRequest(deleteBucketRequest)
	log.Infoln("get DeleteBucket request:", deleteBucketRequest)

	s.bucketRequestCh <- req
	resp := <-req.Done
	bucketInfo := resp.Reply.(*meta.BucketInfo)
	log.Infoln("finish DeleteBucket request:", bucketInfo, ",err:", resp.Err)

	deleteBucketResponse := &pb.DeleteBucketResponse{
		Name:         bucketInfo.Name,
		ObjectsLimit: bucketInfo.UsageInfo.ObjectsLimitCount,
		Capacity:     bucketInfo.UsageInfo.CapacityLimitSize,
		ObjectCount:  bucketInfo.UsageInfo.ObjectsCurrentCount,
	}
	if resp.Err != nil {
		deleteBucketResponse.Message = resp.Err.Error()
	}
	return deleteBucketResponse, resp.Err
}
func (s *BucketService) UpdateBucket(ctx context.Context, updateBucketRequest *pb.UpdateBucketRequest) (*pb.UpdateBucketResponse, error) {
	req := bucket.NewUpdateBucketInfoRequest(updateBucketRequest)
	log.Infoln("get UpdateBucket request:", updateBucketRequest)

	s.bucketRequestCh <- req
	resp := <-req.Done
	bucketInfo := resp.Reply.(*meta.BucketInfo)
	log.Infoln("finish UpdateBucket request:", bucketInfo, ",err:", resp.Err)

	updateBucketResponse := &pb.UpdateBucketResponse{
		Name:         bucketInfo.Name,
		ObjectsLimit: bucketInfo.UsageInfo.ObjectsLimitCount,
		Capacity:     bucketInfo.UsageInfo.CapacityLimitSize,
		ObjectCount:  bucketInfo.UsageInfo.ObjectsCurrentCount,
		BucketDir:    bucketInfo.RealDirName,
	}
	updateBucketResponse.Message = "success"
	if resp.Err != nil {
		updateBucketResponse.Message = resp.Err.Error()
	}
	return updateBucketResponse, resp.Err
}
