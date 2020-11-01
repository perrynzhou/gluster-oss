package bucket

import (
	"github.com/google/uuid"
	meta "gluster-storage-gateway/meta"
	"gluster-storage-gateway/protocol/pb"
)

const (
	CreateBucketType = 1
	DeleteBucketType = 2
	UpdateBucketType = 3
)
const (
	BucketActiveStatus=1
	BucketInActiveStatus=2

)

type BucketInfoRequest struct {
	RequestType uint8
	Info *meta.BucketInfo
	Done chan *BucketInfoResponse
}
type BucketInfoResponse struct {
	Reply interface{}
	Err  error
}

func NewCreateBucketInfoRequest(request *pb.CreateBucketRequest) *BucketInfoRequest {
	return &BucketInfoRequest{
		Info: &meta.BucketInfo{
			Name: request.Name,
			UsageInfo: &meta.BucketUsageInfo{
				ObjectsLimitCount:   request.ObjectsLimit,
				ObjectsCurrentCount: uint64(0),
				CapacityLimitSize: request.Capacity,
			},
			Status:BucketActiveStatus,
			RealDirName: uuid.New().String(),
		},
		Done: make(chan *BucketInfoResponse),
		RequestType:CreateBucketType,
	}
}
func NewDeleteBucketInfoRequest(request *pb.DeleteBucketRequest) *BucketInfoRequest {
	return &BucketInfoRequest{
		Info: &meta.BucketInfo{
			Name: request.Name,
		},
		Done: make(chan *BucketInfoResponse),
		RequestType:DeleteBucketType,
	}
}
func NewUpdateBucketInfoRequest(request *pb.UpdateBucketRequest) *BucketInfoRequest {
	return &BucketInfoRequest{
		Info: &meta.BucketInfo{
			Name: request.Name,
			UsageInfo: &meta.BucketUsageInfo{
				ObjectsLimitCount:   request.ObjectsLimit,
				CapacityLimitSize: request.Capacity,
			},
			Status:BucketActiveStatus,
		},
		Done: make(chan *BucketInfoResponse),
		RequestType:UpdateBucketType,
	}
}
