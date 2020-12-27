package manage

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
	ObjectServiceName = "object"
)

type ObjectRequest struct {
	RequestType uint8
	Info        *meta.ObjectInfo
	Done        chan *ObjectResponse
	LocalFile   string
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
