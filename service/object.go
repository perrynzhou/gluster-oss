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

type ObjectService struct {
	fsApi           *fs_api.FsApi
	ServiceName     string
	objectRequestCh chan *manage.ObjectRequest
	objectMange     *manage.ObjectManage
	wg              *sync.WaitGroup
}

func NewObjectSerivce(api *fs_api.FsApi, serviceName string, bucketCache map[string]*manage.Bucket, notifyBucketCh chan *manage.Bucket, wg *sync.WaitGroup) *ObjectService {
	var err error
	objectService := &ObjectService{
		fsApi:           api,
		ServiceName:     serviceName,
		objectRequestCh: make(chan *manage.ObjectRequest),
		wg:              wg,
	}
	objectService.objectMange = manage.NewObjectManage(api, bucketCache, notifyBucketCh, objectService.objectRequestCh, wg)
	if err != nil {
		log.Errorln("new NewObjectSerivce failed")
		return nil
	}
	log.Info("init ObjectService success")
	return objectService
}

func (s *ObjectService) Run() {
	s.objectMange.Run()
}
func (s *ObjectService) Stop() {
	s.objectMange.Stop()
}
func (s *ObjectService) PutObject(ctx context.Context, putObjectRequest *pb.PutObjectRequest) (*pb.PutObjectResponse, error) {
	req := manage.NewPutObjectRequest(putObjectRequest)
	log.Infoln("get putObjectRequest request:", putObjectRequest)
	s.objectRequestCh <- req
	resp := <-req.Done

	objectInfo := resp.Reply.(*meta.ObjectInfo)
	log.Infoln("finish putObject request:", objectInfo, ",err:", resp.Err)
	putObjectResponse := &pb.PutObjectResponse{
		Message: "success",
	}
	if resp.Err != nil {
		putObjectResponse.Message = resp.Err.Error()
	}
	return putObjectResponse, resp.Err
}
func (s *ObjectService) GetObject(ctx context.Context, getObjectRequest *pb.GetObjectRequest) (*pb.GetObjectReponse, error) {
	req := manage.NewGetbjectRequest(getObjectRequest)
	log.Infoln("get putObjectRequest request:", getObjectRequest)
	s.objectRequestCh <- req
	resp := <-req.Done

	objectInfo := resp.Reply.(*meta.ObjectInfo)
	log.Infoln("finish putObject request:", objectInfo, ",err:", resp.Err)
	getObjectResponse := &pb.GetObjectReponse{
		Message: "success",
	}
	if resp.Err != nil {
		getObjectResponse.Message = resp.Err.Error()
	}
	return getObjectResponse, resp.Err
}
func (s *ObjectService) DeleteObject(ctx context.Context, delObjectRequest *pb.DeleteObjectRequest) (*pb.DeleteBucketResponse, error) {
	req := manage.NewDeleteObjectRequest(delObjectRequest)
	log.Infoln("get putObjectRequest request:", delObjectRequest)
	s.objectRequestCh <- req
	resp := <-req.Done

	objectInfo := resp.Reply.(*meta.ObjectInfo)
	log.Infoln("finish putObject request:", objectInfo, ",err:", resp.Err)
	delObjectResponse := &pb.DeleteBucketResponse{
		Message: "success",
	}
	if resp.Err != nil {
		delObjectResponse.Message = resp.Err.Error()
	}
	return delObjectResponse, resp.Err
}

/*
func (s *ObjectService) DeleteBucket(ctx context.Context, deleteBucketRequest *pb.DeleteBucketRequest) (*pb.DeleteBucketResponse, error) {
	req := manage.NewDeleteBucketRequest(deleteBucketRequest)
	log.Infoln("get DeleteBucket request:", deleteBucketRequest)

	s.bucketRequestCh <- req
	resp := <-req.Done
	bucketInfo := resp.Reply.(*meta.BucketInfo)
	log.Infoln("finish DeleteBucket request:", bucketInfo, ",err:", resp.Err)


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
func (s *ObjectService) UpdateBucket(ctx context.Context, updateBucketRequest *pb.UpdateBucketRequest) (*pb.UpdateBucketResponse, error) {
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
*/
