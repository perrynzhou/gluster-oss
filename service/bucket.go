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
		//request storage LimitBytes
		MaxStorageBytes: bucketInfo.MaxStorageBytes,
		//obejcts limits
		MaxObjectCount: bucketInfo.MaxObjectCount,
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

	/*
	type DeleteBucketResponse struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields

		Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
		BucketDir    string `protobuf:"bytes,2,opt,name=bucket_dir,json=bucketDir,proto3" json:"bucket_dir,omitempty"`
		Key          string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
		CurrentStorageBytes uint64 `protobuf:"varint,4,opt,name=current_bytes,json=CurrentStorageBytes,proto3" json:"current_bytes,omitempty"`
		//obejcts limits
		CurrentObjectCount uint64 `protobuf:"varint,5,opt,name=current_count,json=CurrentObjectCount,proto3" json:"current_count,omitempty"`
		Message      string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	}
	 */
	deleteBucketResponse := &pb.DeleteBucketResponse{
		Name:         bucketInfo.Name,
		BucketDir: bucketInfo.RealDirName,
		CurrentStorageBytes:     bucketInfo.CurrentStorageBytes,
		CurrentObjectCount:  bucketInfo.CurrentObjectCount,
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
		MaxObjectCount: bucketInfo.MaxObjectCount,
		MaxStorageBytes:     bucketInfo.MaxStorageBytes,
		BucketDir:    bucketInfo.RealDirName,
	}
	updateBucketResponse.Message = "success"
	if resp.Err != nil {
		updateBucketResponse.Message = resp.Err.Error()
	}
	return updateBucketResponse, resp.Err
}
