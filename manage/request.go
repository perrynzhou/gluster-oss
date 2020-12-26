package manage

import (
	"fmt"
	meta "glusterfs-storage-gateway/meta"
	"glusterfs-storage-gateway/protocol/pb"

	"github.com/google/uuid"
)

const (
	CreateBucketType = 1
	DeleteBucketType = 2
	UpdateBucketType = 3
)
const (
	ServiceName = "bucket"
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
			LimitCount:   request.ObjectsLimit,
			CurrentCount: uint64(0),
			LimitSize:    request.Capacity,
			Status:       meta.ActiveBucket,
			RealDirName:  fmt.Sprintf("%s-%s", request.Name, uuid.New().String()),
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
			LimitCount: request.ObjectsLimit,
			LimitSize:  request.Capacity,
			Status:     meta.ActiveBucket,
		},
		Done:        make(chan *BucketResponse),
		RequestType: UpdateBucketType,
	}
}
