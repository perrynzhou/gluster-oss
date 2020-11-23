package service

import (
	fs_api "glusterfs-storage-gateway/fs-api"
	"glusterfs-storage-gateway/manage/bucket"
	"glusterfs-storage-gateway/manage/object"
	"glusterfs-storage-gateway/protocol/pb"
	"sync"

	"golang.org/x/net/context"
)

type ObjectService struct {
	fsApi           *fs_api.FsApi
	ServiceName     string
	objectRequestCh chan *object.ObjectInfoRequest
	objectMange     *bucket.BucketManage
	wg              *sync.WaitGroup
}

func NewObjectSerivce(api *fs_api.FsApi, serviceName string, wg *sync.WaitGroup) *ObjectService {
	objectService := &ObjectService{
		fsApi:           api,
		ServiceName:     serviceName,
		objectRequestCh: make(chan *object.ObjectInfoRequest),
		wg:              wg,
	}
	return objectService
}

func (s *ObjectService) Run() {
	s.objectMange.Run()
}
func (s *ObjectService) Stop() {
	s.objectMange.Stop()
}
func (s *ObjectService) PutObject(ctx context.Context, putObjectRequest *pb.PutObjectRequest) (*pb.PutObjectResponse, error) {
	return nil, nil
}
func (s *ObjectService) GetObject(ctx context.Context, getObjectRequest *pb.GetObjectRequest) (*pb.GetObjectReponse, error) {
	return nil, nil
}

func (s *ObjectService) DelObject(ctx context.Context, delObjectRequest *pb.DeleteObjectRequest) (*pb.DeleteObjectResponse, error) {
	return nil, nil
}
