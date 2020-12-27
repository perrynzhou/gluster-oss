package manage

import (
	meta "glusterfs-storage-gateway/meta"
	"glusterfs-storage-gateway/protocol/pb"
)

const (
	CreateBucketType = 1
	DeleteBucketType = 2
	UpdateBucketType = 3
)
const (
	BucketServiceName = "bucket"
)

type BucketRequest struct {
	RequestType uint8
	Info        *meta.BucketInfo
	Done        chan *BucketResponse
}
type BucketResponse struct {
	Reply interface{}
	Err   error
}


func NewCreateBucketRequest(request *pb.CreateBucketRequest) *BucketRequest {
	return &BucketRequest{
		Info: &meta.BucketInfo{
			Name:         request.Name,
			MaxObjectCount:   request.MaxObjectCount,
			CurrentObjectCount: int64(0),
			MaxStorageBytes:    request.MaxStorageBytes,
			Status:       meta.ActiveBucket,
		},
		Done:        make(chan *BucketResponse),
		RequestType: CreateBucketType,
	}
}
func NewDeleteBucketRequest(request *pb.DeleteBucketRequest) *BucketRequest {
	return &BucketRequest{
		Info: &meta.BucketInfo{
			Name:   request.Name,
			Status: meta.InactiveBucket,
		},
		Done:        make(chan *BucketResponse),
		RequestType: DeleteBucketType,
	}
}
func NewUpdateBucketRequest(request *pb.UpdateBucketRequest) *BucketRequest {
	return &BucketRequest{
		Info: &meta.BucketInfo{
			Name:       request.Name,
			MaxObjectCount: request.MaxObjectCount,
			MaxStorageBytes:  request.MaxStorageBytes,
			Status:     meta.ActiveBucket,
		},
		Done:        make(chan *BucketResponse),
		RequestType: UpdateBucketType,
	}
}
