package object

import (
	meta "glusterfs-storage-gateway/meta"
	"glusterfs-storage-gateway/protocol/pb"
)

const (
	PutObjectType = 1
	GetObjectType = 2
	DelObjectType = 3
)
const (
	ServiceName = "object"
)
const (
	ObjectActiveStatus   = 1
	ObjectInActiveStatus = 2
)

type ObjectInfoRequest struct {
	RequestType uint8
	Info        *meta.ObjectInfo
	Done        chan *ObjectInfoResponse
}
type ObjectInfoResponse struct {
	Reply interface{}
	Err   error
}

func NewPutObjectInfoRequest(request *pb.PutObjectRequest) *ObjectInfoRequest {
	return &ObjectInfoRequest{
		RequestType:PutObjectType,
		Info:&meta.ObjectInfo{
			// Name of the bucket.
			Bucket:request.BucketName,
			Name:request.ObjectName,
			Size:request.ObjectsSize,
			ContentType:request.ContentType,
			UserTags:request.UserTags,
		},
	}
}
func NewDeleteObjectInfoRequest(request *pb.DeleteObjectRequest) *ObjectInfoRequest {
	return &ObjectInfoRequest{
		RequestType:DelObjectType,
		Info:&meta.ObjectInfo{
			Id:request.ObjectId,
			Bucket: request.BucketName,
			Status:ObjectInActiveStatus,
		},
	}
}
func NewGetObjectInfoRequest(request *pb.GetObjectRequest) *ObjectInfoRequest {
	return &ObjectInfoRequest{
		RequestType:GetObjectType,
		Info:&meta.ObjectInfo{
			Id:request.ObjectId,
			Bucket:request.BucketName,
			Status:ObjectActiveStatus,
		},
	}
}
