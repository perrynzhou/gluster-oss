// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: service.proto

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

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbf, 0x04,
	0x0a, 0x15, 0x47, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x5e, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x5e, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x5e, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x52, 0x0a, 0x09, 0x50, 0x75, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x74, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x5e, 0x0a, 0x0c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x52, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x01, 0x2a, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_proto_goTypes = []interface{}{
	(*CreateBucketRequest)(nil),  // 0: pb.CreateBucketRequest
	(*DeleteBucketRequest)(nil),  // 1: pb.DeleteBucketRequest
	(*UpdateBucketRequest)(nil),  // 2: pb.UpdateBucketRequest
	(*PutObjectRequest)(nil),     // 3: pb.PutObjectRequest
	(*DeleteObjectRequest)(nil),  // 4: pb.DeleteObjectRequest
	(*GetObjectRequest)(nil),     // 5: pb.GetObjectRequest
	(*CreateBucketResponse)(nil), // 6: pb.CreateBucketResponse
	(*DeleteBucketResponse)(nil), // 7: pb.DeleteBucketResponse
	(*UpdateBucketResponse)(nil), // 8: pb.UpdateBucketResponse
	(*PutObjectResponse)(nil),    // 9: pb.PutObjectResponse
	(*DeleteObjectResponse)(nil), // 10: pb.DeleteObjectResponse
	(*GetObjectResponse)(nil),    // 11: pb.GetObjectResponse
}
var file_service_proto_depIdxs = []int32{
	0,  // 0: pb.GlusterStorageGateway.CreateBucket:input_type -> pb.CreateBucketRequest
	1,  // 1: pb.GlusterStorageGateway.DeleteBucket:input_type -> pb.DeleteBucketRequest
	2,  // 2: pb.GlusterStorageGateway.UpdateBucket:input_type -> pb.UpdateBucketRequest
	3,  // 3: pb.GlusterStorageGateway.PutObject:input_type -> pb.PutObjectRequest
	4,  // 4: pb.GlusterStorageGateway.DeleteObject:input_type -> pb.DeleteObjectRequest
	5,  // 5: pb.GlusterStorageGateway.GetObject:input_type -> pb.GetObjectRequest
	6,  // 6: pb.GlusterStorageGateway.CreateBucket:output_type -> pb.CreateBucketResponse
	7,  // 7: pb.GlusterStorageGateway.DeleteBucket:output_type -> pb.DeleteBucketResponse
	8,  // 8: pb.GlusterStorageGateway.UpdateBucket:output_type -> pb.UpdateBucketResponse
	9,  // 9: pb.GlusterStorageGateway.PutObject:output_type -> pb.PutObjectResponse
	10, // 10: pb.GlusterStorageGateway.DeleteObject:output_type -> pb.DeleteObjectResponse
	11, // 11: pb.GlusterStorageGateway.GetObject:output_type -> pb.GetObjectResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_bucket_proto_init()
	file_object_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GlusterStorageGatewayClient is the client API for GlusterStorageGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GlusterStorageGatewayClient interface {
	CreateBucket(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*CreateBucketResponse, error)
	DeleteBucket(ctx context.Context, in *DeleteBucketRequest, opts ...grpc.CallOption) (*DeleteBucketResponse, error)
	UpdateBucket(ctx context.Context, in *UpdateBucketRequest, opts ...grpc.CallOption) (*UpdateBucketResponse, error)
	PutObject(ctx context.Context, in *PutObjectRequest, opts ...grpc.CallOption) (*PutObjectResponse, error)
	DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectResponse, error)
	GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectResponse, error)
}

type glusterStorageGatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGlusterStorageGatewayClient(cc grpc.ClientConnInterface) GlusterStorageGatewayClient {
	return &glusterStorageGatewayClient{cc}
}

func (c *glusterStorageGatewayClient) CreateBucket(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*CreateBucketResponse, error) {
	out := new(CreateBucketResponse)
	err := c.cc.Invoke(ctx, "/pb.GlusterStorageGateway/CreateBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *glusterStorageGatewayClient) DeleteBucket(ctx context.Context, in *DeleteBucketRequest, opts ...grpc.CallOption) (*DeleteBucketResponse, error) {
	out := new(DeleteBucketResponse)
	err := c.cc.Invoke(ctx, "/pb.GlusterStorageGateway/DeleteBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *glusterStorageGatewayClient) UpdateBucket(ctx context.Context, in *UpdateBucketRequest, opts ...grpc.CallOption) (*UpdateBucketResponse, error) {
	out := new(UpdateBucketResponse)
	err := c.cc.Invoke(ctx, "/pb.GlusterStorageGateway/UpdateBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *glusterStorageGatewayClient) PutObject(ctx context.Context, in *PutObjectRequest, opts ...grpc.CallOption) (*PutObjectResponse, error) {
	out := new(PutObjectResponse)
	err := c.cc.Invoke(ctx, "/pb.GlusterStorageGateway/PutObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *glusterStorageGatewayClient) DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectResponse, error) {
	out := new(DeleteObjectResponse)
	err := c.cc.Invoke(ctx, "/pb.GlusterStorageGateway/DeleteObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *glusterStorageGatewayClient) GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectResponse, error) {
	out := new(GetObjectResponse)
	err := c.cc.Invoke(ctx, "/pb.GlusterStorageGateway/GetObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GlusterStorageGatewayServer is the server API for GlusterStorageGateway service.
type GlusterStorageGatewayServer interface {
	CreateBucket(context.Context, *CreateBucketRequest) (*CreateBucketResponse, error)
	DeleteBucket(context.Context, *DeleteBucketRequest) (*DeleteBucketResponse, error)
	UpdateBucket(context.Context, *UpdateBucketRequest) (*UpdateBucketResponse, error)
	PutObject(context.Context, *PutObjectRequest) (*PutObjectResponse, error)
	DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error)
	GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error)
}

// UnimplementedGlusterStorageGatewayServer can be embedded to have forward compatible implementations.
type UnimplementedGlusterStorageGatewayServer struct {
}

func (*UnimplementedGlusterStorageGatewayServer) CreateBucket(context.Context, *CreateBucketRequest) (*CreateBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBucket not implemented")
}
func (*UnimplementedGlusterStorageGatewayServer) DeleteBucket(context.Context, *DeleteBucketRequest) (*DeleteBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBucket not implemented")
}
func (*UnimplementedGlusterStorageGatewayServer) UpdateBucket(context.Context, *UpdateBucketRequest) (*UpdateBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBucket not implemented")
}
func (*UnimplementedGlusterStorageGatewayServer) PutObject(context.Context, *PutObjectRequest) (*PutObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutObject not implemented")
}
func (*UnimplementedGlusterStorageGatewayServer) DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObject not implemented")
}
func (*UnimplementedGlusterStorageGatewayServer) GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObject not implemented")
}

func RegisterGlusterStorageGatewayServer(s *grpc.Server, srv GlusterStorageGatewayServer) {
	s.RegisterService(&_GlusterStorageGateway_serviceDesc, srv)
}

func _GlusterStorageGateway_CreateBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlusterStorageGatewayServer).CreateBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GlusterStorageGateway/CreateBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlusterStorageGatewayServer).CreateBucket(ctx, req.(*CreateBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlusterStorageGateway_DeleteBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlusterStorageGatewayServer).DeleteBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GlusterStorageGateway/DeleteBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlusterStorageGatewayServer).DeleteBucket(ctx, req.(*DeleteBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlusterStorageGateway_UpdateBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlusterStorageGatewayServer).UpdateBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GlusterStorageGateway/UpdateBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlusterStorageGatewayServer).UpdateBucket(ctx, req.(*UpdateBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlusterStorageGateway_PutObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlusterStorageGatewayServer).PutObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GlusterStorageGateway/PutObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlusterStorageGatewayServer).PutObject(ctx, req.(*PutObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlusterStorageGateway_DeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlusterStorageGatewayServer).DeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GlusterStorageGateway/DeleteObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlusterStorageGatewayServer).DeleteObject(ctx, req.(*DeleteObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlusterStorageGateway_GetObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlusterStorageGatewayServer).GetObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GlusterStorageGateway/GetObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlusterStorageGatewayServer).GetObject(ctx, req.(*GetObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GlusterStorageGateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GlusterStorageGateway",
	HandlerType: (*GlusterStorageGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBucket",
			Handler:    _GlusterStorageGateway_CreateBucket_Handler,
		},
		{
			MethodName: "DeleteBucket",
			Handler:    _GlusterStorageGateway_DeleteBucket_Handler,
		},
		{
			MethodName: "UpdateBucket",
			Handler:    _GlusterStorageGateway_UpdateBucket_Handler,
		},
		{
			MethodName: "PutObject",
			Handler:    _GlusterStorageGateway_PutObject_Handler,
		},
		{
			MethodName: "DeleteObject",
			Handler:    _GlusterStorageGateway_DeleteObject_Handler,
		},
		{
			MethodName: "GetObject",
			Handler:    _GlusterStorageGateway_GetObject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
