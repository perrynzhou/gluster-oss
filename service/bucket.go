package service

import (
	fs_api "glusterfs-storage-gateway/fs-api"
	"glusterfs-storage-gateway/manage"
	"glusterfs-storage-gateway/meta"
	"glusterfs-storage-gateway/protocol/pb"
	"sync"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)
const (
	BucketObjectMetaSubffix  = ".obj.meta"
)
type BucketService struct {
	fsApi           *fs_api.FsApi
	ServiceName     string
	bucketRequestCh chan *manage.BucketRequest
	bucketMange     *manage.BucketManage
	wg              *sync.WaitGroup
}

func NewBucketSerivce(api *fs_api.FsApi, serviceName string, wg *sync.WaitGroup) *BucketService {
	var err error
	bucketService := &BucketService{
		fsApi:           api,
		ServiceName:     serviceName,
		bucketRequestCh: make(chan *manage.BucketRequest),
		wg:              wg,
	}
	bucketService.bucketMange, err = manage.NewBucketManage(api, bucketService.bucketRequestCh, wg)
	if err != nil {
		log.Errorln("new NewBucketManage failed")
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
	req := manage.NewCreateBucketRequest(createBucketRequest)
	log.Infoln("get CreateBucket request:", createBucketRequest)
	s.bucketRequestCh <- req
	resp := <-req.Done

	bucketInfo := resp.Reply.(*meta.BucketInfo)
	log.Infoln("finish CreateBucket request:", bucketInfo, ",err:", resp.Err)
	createBucketResponse := &pb.CreateBucketResponse{
		Name: bucketInfo.Name,
		//request storage capacity
		Capacity: bucketInfo.LimitSize,
		//obejcts limits
		ObjectsLimit: bucketInfo.LimitCount,
		BucketDir:    bucketInfo.RealDirName,
		Message:      "success",
	}
	if resp.Err != nil {
		createBucketResponse.Message = resp.Err.Error()
	}
	return createBucketResponse, resp.Err
}

func (s *BucketService) DeleteBucket(ctx context.Context, deleteBucketRequest *pb.DeleteBucketRequest) (*pb.DeleteBucketResponse, error) {
	req := manage.NewDeleteBucketRequest(deleteBucketRequest)
	log.Infoln("get DeleteBucket request:", deleteBucketRequest)

	s.bucketRequestCh <- req
	resp := <-req.Done
	bucketInfo := resp.Reply.(*meta.BucketInfo)
	log.Infoln("finish DeleteBucket request:", bucketInfo, ",err:", resp.Err)

	deleteBucketResponse := &pb.DeleteBucketResponse{
		Name:         bucketInfo.Name,
		ObjectsLimit: bucketInfo.LimitCount,
		Capacity:     bucketInfo.LimitSize,
		ObjectCount:  bucketInfo.CurrentCount,
	}
	if resp.Err != nil {
		deleteBucketResponse.Message = resp.Err.Error()
	}
	return deleteBucketResponse, resp.Err
}
func (s *BucketService) UpdateBucket(ctx context.Context, updateBucketRequest *pb.UpdateBucketRequest) (*pb.UpdateBucketResponse, error) {
	req := manage.NewUpdateBucketRequest(updateBucketRequest)
	log.Infoln("get UpdateBucket request:", updateBucketRequest)

	s.bucketRequestCh <- req
	resp := <-req.Done
	bucketInfo := resp.Reply.(*meta.BucketInfo)
	log.Infoln("finish UpdateBucket request:", bucketInfo, ",err:", resp.Err)

	updateBucketResponse := &pb.UpdateBucketResponse{
		Name:         bucketInfo.Name,
		ObjectsLimit: bucketInfo.LimitCount,
		Capacity:     bucketInfo.LimitSize,
		ObjectCount:  bucketInfo.CurrentCount,
		BucketDir:    bucketInfo.RealDirName,
	}
	updateBucketResponse.Message = "success"
	if resp.Err != nil {
		updateBucketResponse.Message = resp.Err.Error()
	}
	return updateBucketResponse, resp.Err
}
