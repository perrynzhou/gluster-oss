// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: bucket.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateBucketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//request bucket name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//access privileges
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	//request storage capacity
	Capacity uint64 `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	//obejcts limits
	ObjectsLimit uint64 `protobuf:"varint,4,opt,name=objects_limit,json=objectsLimit,proto3" json:"objects_limit,omitempty"`
}

func (x *CreateBucketRequest) Reset() {
	*x = CreateBucketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBucketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBucketRequest) ProtoMessage() {}

func (x *CreateBucketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBucketRequest.ProtoReflect.Descriptor instead.
func (*CreateBucketRequest) Descriptor() ([]byte, []int) {
	return file_bucket_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBucketRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBucketRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CreateBucketRequest) GetCapacity() uint64 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *CreateBucketRequest) GetObjectsLimit() uint64 {
	if x != nil {
		return x.ObjectsLimit
	}
	return 0
}

type CreateBucketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requst  *CreateBucketRequest `protobuf:"bytes,1,opt,name=requst,proto3" json:"requst,omitempty"`
	Message string               `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateBucketResponse) Reset() {
	*x = CreateBucketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBucketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBucketResponse) ProtoMessage() {}

func (x *CreateBucketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBucketResponse.ProtoReflect.Descriptor instead.
func (*CreateBucketResponse) Descriptor() ([]byte, []int) {
	return file_bucket_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBucketResponse) GetRequst() *CreateBucketRequest {
	if x != nil {
		return x.Requst
	}
	return nil
}

func (x *CreateBucketResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteBucketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//bucket name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteBucketRequest) Reset() {
	*x = DeleteBucketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBucketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBucketRequest) ProtoMessage() {}

func (x *DeleteBucketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBucketRequest.ProtoReflect.Descriptor instead.
func (*DeleteBucketRequest) Descriptor() ([]byte, []int) {
	return file_bucket_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteBucketRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteBucketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Key          string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Capacity     uint64 `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	ObjectsLimit uint64 `protobuf:"varint,4,opt,name=objects_limit,json=objectsLimit,proto3" json:"objects_limit,omitempty"`
	ObjectCount  uint64 `protobuf:"varint,5,opt,name=object_count,json=objectCount,proto3" json:"object_count,omitempty"`
	BucketDir    string `protobuf:"bytes,6,opt,name=bucket_dir,json=bucketDir,proto3" json:"bucket_dir,omitempty"`
	Message      string `protobuf:"bytes,7,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteBucketResponse) Reset() {
	*x = DeleteBucketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBucketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBucketResponse) ProtoMessage() {}

func (x *DeleteBucketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBucketResponse.ProtoReflect.Descriptor instead.
func (*DeleteBucketResponse) Descriptor() ([]byte, []int) {
	return file_bucket_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteBucketResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeleteBucketResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *DeleteBucketResponse) GetCapacity() uint64 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *DeleteBucketResponse) GetObjectsLimit() uint64 {
	if x != nil {
		return x.ObjectsLimit
	}
	return 0
}

func (x *DeleteBucketResponse) GetObjectCount() uint64 {
	if x != nil {
		return x.ObjectCount
	}
	return 0
}

func (x *DeleteBucketResponse) GetBucketDir() string {
	if x != nil {
		return x.BucketDir
	}
	return ""
}

func (x *DeleteBucketResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UpdateBucketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//request bucket name
	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Key          string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Capacity     uint64 `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	ObjectsLimit uint64 `protobuf:"varint,4,opt,name=objects_limit,json=objectsLimit,proto3" json:"objects_limit,omitempty"`
}

func (x *UpdateBucketRequest) Reset() {
	*x = UpdateBucketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBucketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBucketRequest) ProtoMessage() {}

func (x *UpdateBucketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBucketRequest.ProtoReflect.Descriptor instead.
func (*UpdateBucketRequest) Descriptor() ([]byte, []int) {
	return file_bucket_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateBucketRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBucketRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UpdateBucketRequest) GetCapacity() uint64 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *UpdateBucketRequest) GetObjectsLimit() uint64 {
	if x != nil {
		return x.ObjectsLimit
	}
	return 0
}

type UpdateBucketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Capacity     uint64 `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	ObjectsLimit uint64 `protobuf:"varint,3,opt,name=objects_limit,json=objectsLimit,proto3" json:"objects_limit,omitempty"`
	ObjectCount  uint64 `protobuf:"varint,4,opt,name=object_count,json=objectCount,proto3" json:"object_count,omitempty"`
	BucketDir    string `protobuf:"bytes,5,opt,name=bucket_dir,json=bucketDir,proto3" json:"bucket_dir,omitempty"`
	Message      string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateBucketResponse) Reset() {
	*x = UpdateBucketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bucket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBucketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBucketResponse) ProtoMessage() {}

func (x *UpdateBucketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bucket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBucketResponse.ProtoReflect.Descriptor instead.
func (*UpdateBucketResponse) Descriptor() ([]byte, []int) {
	return file_bucket_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBucketResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBucketResponse) GetCapacity() uint64 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *UpdateBucketResponse) GetObjectsLimit() uint64 {
	if x != nil {
		return x.ObjectsLimit
	}
	return 0
}

func (x *UpdateBucketResponse) GetObjectCount() uint64 {
	if x != nil {
		return x.ObjectCount
	}
	return 0
}

func (x *UpdateBucketResponse) GetBucketDir() string {
	if x != nil {
		return x.BucketDir
	}
	return ""
}

func (x *UpdateBucketResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_bucket_proto protoreflect.FileDescriptor

var file_bucket_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x7c, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x61,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x75, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x71, 0x75, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x29, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xd9, 0x01, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x44, 0x69, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x7c, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xc7, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x23, 0x0a, 0x0d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x44, 0x69, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0xbd, 0x02, 0x0a, 0x1b, 0x47, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x5e, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x3a, 0x01, 0x2a,
	0x12, 0x5e, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x3a, 0x01, 0x2a,
	0x12, 0x5e, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x3a, 0x01, 0x2a,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bucket_proto_rawDescOnce sync.Once
	file_bucket_proto_rawDescData = file_bucket_proto_rawDesc
)

func file_bucket_proto_rawDescGZIP() []byte {
	file_bucket_proto_rawDescOnce.Do(func() {
		file_bucket_proto_rawDescData = protoimpl.X.CompressGZIP(file_bucket_proto_rawDescData)
	})
	return file_bucket_proto_rawDescData
}

var file_bucket_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bucket_proto_goTypes = []interface{}{
	(*CreateBucketRequest)(nil),  // 0: pb.CreateBucketRequest
	(*CreateBucketResponse)(nil), // 1: pb.CreateBucketResponse
	(*DeleteBucketRequest)(nil),  // 2: pb.DeleteBucketRequest
	(*DeleteBucketResponse)(nil), // 3: pb.DeleteBucketResponse
	(*UpdateBucketRequest)(nil),  // 4: pb.UpdateBucketRequest
	(*UpdateBucketResponse)(nil), // 5: pb.UpdateBucketResponse
}
var file_bucket_proto_depIdxs = []int32{
	0, // 0: pb.CreateBucketResponse.requst:type_name -> pb.CreateBucketRequest
	0, // 1: pb.GlusterStorageGatewayBucket.CreateBucket:input_type -> pb.CreateBucketRequest
	2, // 2: pb.GlusterStorageGatewayBucket.DeleteBucket:input_type -> pb.DeleteBucketRequest
	4, // 3: pb.GlusterStorageGatewayBucket.UpdateBucket:input_type -> pb.UpdateBucketRequest
	1, // 4: pb.GlusterStorageGatewayBucket.CreateBucket:output_type -> pb.CreateBucketResponse
	3, // 5: pb.GlusterStorageGatewayBucket.DeleteBucket:output_type -> pb.DeleteBucketResponse
	5, // 6: pb.GlusterStorageGatewayBucket.UpdateBucket:output_type -> pb.UpdateBucketResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bucket_proto_init() }
func file_bucket_proto_init() {
	if File_bucket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bucket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBucketRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bucket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBucketResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bucket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBucketRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bucket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBucketResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bucket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBucketRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bucket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBucketResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bucket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bucket_proto_goTypes,
		DependencyIndexes: file_bucket_proto_depIdxs,
		MessageInfos:      file_bucket_proto_msgTypes,
	}.Build()
	File_bucket_proto = out.File
	file_bucket_proto_rawDesc = nil
	file_bucket_proto_goTypes = nil
	file_bucket_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GlusterStorageGatewayBucketClient is the client API for GlusterStorageGatewayBucket service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GlusterStorageGatewayBucketClient interface {
	CreateBucket(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*CreateBucketResponse, error)
	DeleteBucket(ctx context.Context, in *DeleteBucketRequest, opts ...grpc.CallOption) (*DeleteBucketResponse, error)
	UpdateBucket(ctx context.Context, in *UpdateBucketRequest, opts ...grpc.CallOption) (*UpdateBucketResponse, error)
}

type glusterStorageGatewayBucketClient struct {
	cc grpc.ClientConnInterface
}

func NewGlusterStorageGatewayBucketClient(cc grpc.ClientConnInterface) GlusterStorageGatewayBucketClient {
	return &glusterStorageGatewayBucketClient{cc}
}

func (c *glusterStorageGatewayBucketClient) CreateBucket(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*CreateBucketResponse, error) {
	out := new(CreateBucketResponse)
	err := c.cc.Invoke(ctx, "/pb.GlusterStorageGatewayBucket/CreateBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *glusterStorageGatewayBucketClient) DeleteBucket(ctx context.Context, in *DeleteBucketRequest, opts ...grpc.CallOption) (*DeleteBucketResponse, error) {
	out := new(DeleteBucketResponse)
	err := c.cc.Invoke(ctx, "/pb.GlusterStorageGatewayBucket/DeleteBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *glusterStorageGatewayBucketClient) UpdateBucket(ctx context.Context, in *UpdateBucketRequest, opts ...grpc.CallOption) (*UpdateBucketResponse, error) {
	out := new(UpdateBucketResponse)
	err := c.cc.Invoke(ctx, "/pb.GlusterStorageGatewayBucket/UpdateBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GlusterStorageGatewayBucketServer is the server API for GlusterStorageGatewayBucket service.
type GlusterStorageGatewayBucketServer interface {
	CreateBucket(context.Context, *CreateBucketRequest) (*CreateBucketResponse, error)
	DeleteBucket(context.Context, *DeleteBucketRequest) (*DeleteBucketResponse, error)
	UpdateBucket(context.Context, *UpdateBucketRequest) (*UpdateBucketResponse, error)
}

// UnimplementedGlusterStorageGatewayBucketServer can be embedded to have forward compatible implementations.
type UnimplementedGlusterStorageGatewayBucketServer struct {
}

func (*UnimplementedGlusterStorageGatewayBucketServer) CreateBucket(context.Context, *CreateBucketRequest) (*CreateBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBucket not implemented")
}
func (*UnimplementedGlusterStorageGatewayBucketServer) DeleteBucket(context.Context, *DeleteBucketRequest) (*DeleteBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBucket not implemented")
}
func (*UnimplementedGlusterStorageGatewayBucketServer) UpdateBucket(context.Context, *UpdateBucketRequest) (*UpdateBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBucket not implemented")
}

func RegisterGlusterStorageGatewayBucketServer(s *grpc.Server, srv GlusterStorageGatewayBucketServer) {
	s.RegisterService(&_GlusterStorageGatewayBucket_serviceDesc, srv)
}

func _GlusterStorageGatewayBucket_CreateBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlusterStorageGatewayBucketServer).CreateBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GlusterStorageGatewayBucket/CreateBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlusterStorageGatewayBucketServer).CreateBucket(ctx, req.(*CreateBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlusterStorageGatewayBucket_DeleteBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlusterStorageGatewayBucketServer).DeleteBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GlusterStorageGatewayBucket/DeleteBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlusterStorageGatewayBucketServer).DeleteBucket(ctx, req.(*DeleteBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlusterStorageGatewayBucket_UpdateBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlusterStorageGatewayBucketServer).UpdateBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GlusterStorageGatewayBucket/UpdateBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlusterStorageGatewayBucketServer).UpdateBucket(ctx, req.(*UpdateBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GlusterStorageGatewayBucket_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GlusterStorageGatewayBucket",
	HandlerType: (*GlusterStorageGatewayBucketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBucket",
			Handler:    _GlusterStorageGatewayBucket_CreateBucket_Handler,
		},
		{
			MethodName: "DeleteBucket",
			Handler:    _GlusterStorageGatewayBucket_DeleteBucket_Handler,
		},
		{
			MethodName: "UpdateBucket",
			Handler:    _GlusterStorageGatewayBucket_UpdateBucket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bucket.proto",
}