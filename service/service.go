package service

import (

	"gluster-storage-gateway/protocol/pb"
	"golang.org/x/net/context"
)

type Service struct {
	bucketService *BucketService
}
func NewService(bucketService *BucketService) *Service {
	return &Service {
		bucketService:bucketService,
	}
}
func (s *Service) CreateBucket(ctx context.Context, createBucketRequest *pb.CreateBucketRequest) (*pb.CreateBucketResponse, error){
	return s.bucketService.CreateBucket(ctx,createBucketRequest)
}
func (s *Service) DeleteBucket(ctx context.Context, deleteBucketRequest *pb.DeleteBucketRequest) (*pb.DeleteBucketResponse, error){
	return s.bucketService.DeleteBucket(ctx,deleteBucketRequest)
}
func (s *Service) UpdateBucket(ctx context.Context, updateBucketRequest *pb.UpdateBucketRequest) (*pb.UpdateBucketResponse, error){
	return s.bucketService.UpdateBucket(ctx,updateBucketRequest)

}
func (s *Service)Run() {
	s.bucketService.Run()
}


