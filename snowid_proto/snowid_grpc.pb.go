// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: snowid.proto

package gengrpcsnow

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GRPCSnowServiceClient is the client API for GRPCSnowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GRPCSnowServiceClient interface {
	// 获取雪花ID
	GetSnowID(ctx context.Context, in *SnowIDReq, opts ...grpc.CallOption) (*Resp, error)
	// 批量获取雪花ID
	GetBatchSnowID(ctx context.Context, in *BatchSnowIDReq, opts ...grpc.CallOption) (*BatchResp, error)
}

type gRPCSnowServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGRPCSnowServiceClient(cc grpc.ClientConnInterface) GRPCSnowServiceClient {
	return &gRPCSnowServiceClient{cc}
}

func (c *gRPCSnowServiceClient) GetSnowID(ctx context.Context, in *SnowIDReq, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/gengrpcsnow.GRPCSnowService/GetSnowID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCSnowServiceClient) GetBatchSnowID(ctx context.Context, in *BatchSnowIDReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, "/gengrpcsnow.GRPCSnowService/GetBatchSnowID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCSnowServiceServer is the server API for GRPCSnowService service.
// All implementations must embed UnimplementedGRPCSnowServiceServer
// for forward compatibility
type GRPCSnowServiceServer interface {
	// 获取雪花ID
	GetSnowID(context.Context, *SnowIDReq) (*Resp, error)
	// 批量获取雪花ID
	GetBatchSnowID(context.Context, *BatchSnowIDReq) (*BatchResp, error)
	mustEmbedUnimplementedGRPCSnowServiceServer()
}

// UnimplementedGRPCSnowServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGRPCSnowServiceServer struct {
}

func (UnimplementedGRPCSnowServiceServer) GetSnowID(context.Context, *SnowIDReq) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSnowID not implemented")
}
func (UnimplementedGRPCSnowServiceServer) GetBatchSnowID(context.Context, *BatchSnowIDReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBatchSnowID not implemented")
}
func (UnimplementedGRPCSnowServiceServer) mustEmbedUnimplementedGRPCSnowServiceServer() {}

// UnsafeGRPCSnowServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GRPCSnowServiceServer will
// result in compilation errors.
type UnsafeGRPCSnowServiceServer interface {
	mustEmbedUnimplementedGRPCSnowServiceServer()
}

func RegisterGRPCSnowServiceServer(s grpc.ServiceRegistrar, srv GRPCSnowServiceServer) {
	s.RegisterService(&GRPCSnowService_ServiceDesc, srv)
}

func _GRPCSnowService_GetSnowID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnowIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCSnowServiceServer).GetSnowID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gengrpcsnow.GRPCSnowService/GetSnowID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCSnowServiceServer).GetSnowID(ctx, req.(*SnowIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCSnowService_GetBatchSnowID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchSnowIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCSnowServiceServer).GetBatchSnowID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gengrpcsnow.GRPCSnowService/GetBatchSnowID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCSnowServiceServer).GetBatchSnowID(ctx, req.(*BatchSnowIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GRPCSnowService_ServiceDesc is the grpc.ServiceDesc for GRPCSnowService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GRPCSnowService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gengrpcsnow.GRPCSnowService",
	HandlerType: (*GRPCSnowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSnowID",
			Handler:    _GRPCSnowService_GetSnowID_Handler,
		},
		{
			MethodName: "GetBatchSnowID",
			Handler:    _GRPCSnowService_GetBatchSnowID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "snowid.proto",
}
