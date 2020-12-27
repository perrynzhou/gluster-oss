package manage

import (
	meta "glusterfs-storage-gateway/meta"
	"glusterfs-storage-gateway/protocol/pb"
)

const (
	CreateObjectType = 1
	DeleteObjectType = 2
)
const (
	ObjectServiceName = "object"
)

type ObjectRequest struct {
	RequestType uint8
	Info        *meta.ObjectInfo
	Done        chan *ObjectResponse
}
type ObjectResponse struct {
	Reply interface{}
	Err   error
}


func NewPutObjectRequest(request *pb.PutObjectRequest) *ObjectRequest {
	return nil
}
func NewDeleteObjectRequest(request *pb.DeleteObjectRequest) *ObjectRequest {
	return nil
}
func NewGetbjectRequest(request *pb.GetObjectRequest) *ObjectRequest {
	return nil
}
