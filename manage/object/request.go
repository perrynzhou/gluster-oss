package object

import (
	meta "glusterfs-storage-gateway/meta"

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

func NewPutObjectInfoRequest(request *pb.GetObjectRequest) *ObjectInfoRequest {
	return &ObjectInfoRequest{

	}
}
func NewDeleteObjectInfoRequest(request *pb.DelObjectRequest) *ObjectInfoRequest {
	return &ObjectInfoRequest{}
}
func NewGetObjectInfoRequest(request *pb.GetObjectRequest) *ObjectInfoRequest {
	return &ObjectInfoRequest{}
}
