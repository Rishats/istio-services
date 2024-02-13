// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: api/server_info/server_info.proto

package server_info

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

const (
	ServerInfoService_GetServerInfo_FullMethodName = "/api.server_info.ServerInfoService/GetServerInfo"
)

// ServerInfoServiceClient is the client API for ServerInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerInfoServiceClient interface {
	// Sends a request to get server information
	GetServerInfo(ctx context.Context, in *GetServerInfoRequest, opts ...grpc.CallOption) (*GetServerInfoResponse, error)
}

type serverInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerInfoServiceClient(cc grpc.ClientConnInterface) ServerInfoServiceClient {
	return &serverInfoServiceClient{cc}
}

func (c *serverInfoServiceClient) GetServerInfo(ctx context.Context, in *GetServerInfoRequest, opts ...grpc.CallOption) (*GetServerInfoResponse, error) {
	out := new(GetServerInfoResponse)
	err := c.cc.Invoke(ctx, ServerInfoService_GetServerInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerInfoServiceServer is the server API for ServerInfoService service.
// All implementations must embed UnimplementedServerInfoServiceServer
// for forward compatibility
type ServerInfoServiceServer interface {
	// Sends a request to get server information
	GetServerInfo(context.Context, *GetServerInfoRequest) (*GetServerInfoResponse, error)
	mustEmbedUnimplementedServerInfoServiceServer()
}

// UnimplementedServerInfoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServerInfoServiceServer struct {
}

func (UnimplementedServerInfoServiceServer) GetServerInfo(context.Context, *GetServerInfoRequest) (*GetServerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerInfo not implemented")
}
func (UnimplementedServerInfoServiceServer) mustEmbedUnimplementedServerInfoServiceServer() {}

// UnsafeServerInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerInfoServiceServer will
// result in compilation errors.
type UnsafeServerInfoServiceServer interface {
	mustEmbedUnimplementedServerInfoServiceServer()
}

func RegisterServerInfoServiceServer(s grpc.ServiceRegistrar, srv ServerInfoServiceServer) {
	s.RegisterService(&ServerInfoService_ServiceDesc, srv)
}

func _ServerInfoService_GetServerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerInfoServiceServer).GetServerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerInfoService_GetServerInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerInfoServiceServer).GetServerInfo(ctx, req.(*GetServerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerInfoService_ServiceDesc is the grpc.ServiceDesc for ServerInfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerInfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.server_info.ServerInfoService",
	HandlerType: (*ServerInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServerInfo",
			Handler:    _ServerInfoService_GetServerInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/server_info/server_info.proto",
}