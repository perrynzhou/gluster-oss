package bucket

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
const (
	BucketActiveStatus   = 1
	BucketInActiveStatus = 2
)

type BucketInfoRequest struct {
	RequestType uint8
	Info        *meta.BucketInfo
	Done        chan *BucketInfoResponse
}
type BucketInfoResponse struct {
	Reply interface{}
	Err   error
}

func NewCreateBucketInfoRequest(request *pb.CreateBucketRequest) *BucketInfoRequest {
	return &BucketInfoRequest{
		Info: &meta.BucketInfo{
			Name: request.Name,
			UsageInfo: &meta.BucketUsageInfo{
				ObjectsLimitCount:   request.ObjectsLimit,
				ObjectsCurrentCount: uint64(0),
				CapacityLimitSize:   request.Capacity,
			},
			Status:      BucketActiveStatus,
			RealDirName: fmt.Sprintf("%s-%s", request.Name, uuid.New().String()),
		},
		Done:        make(chan *BucketInfoResponse),
		RequestType: CreateBucketType,
	}
}
func NewDeleteBucketInfoRequest(request *pb.DeleteBucketRequest) *BucketInfoRequest {
	return &BucketInfoRequest{
		Info: &meta.BucketInfo{
			Name: request.Name,
		},
		Done:        make(chan *BucketInfoResponse),
		RequestType: DeleteBucketType,
	}
}
func NewUpdateBucketInfoRequest(request *pb.UpdateBucketRequest) *BucketInfoRequest {
	return &BucketInfoRequest{
		Info: &meta.BucketInfo{
			Name: request.Name,
			UsageInfo: &meta.BucketUsageInfo{
				ObjectsLimitCount: request.ObjectsLimit,
				CapacityLimitSize: request.Capacity,
			},
			Status: BucketActiveStatus,
		},
		Done:        make(chan *BucketInfoResponse),
		RequestType: UpdateBucketType,
	}
}
