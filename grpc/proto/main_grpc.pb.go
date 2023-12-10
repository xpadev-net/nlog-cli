// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: proto/main.proto

package main

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

// LoggingServiceClient is the client API for LoggingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoggingServiceClient interface {
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	AppendLog(ctx context.Context, in *AppendLogRequest, opts ...grpc.CallOption) (*AppendLogResponse, error)
	EndTask(ctx context.Context, in *EndTaskRequest, opts ...grpc.CallOption) (*EndTaskResponse, error)
}

type loggingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoggingServiceClient(cc grpc.ClientConnInterface) LoggingServiceClient {
	return &loggingServiceClient{cc}
}

func (c *loggingServiceClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, "/LoggingService/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceClient) AppendLog(ctx context.Context, in *AppendLogRequest, opts ...grpc.CallOption) (*AppendLogResponse, error) {
	out := new(AppendLogResponse)
	err := c.cc.Invoke(ctx, "/LoggingService/AppendLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceClient) EndTask(ctx context.Context, in *EndTaskRequest, opts ...grpc.CallOption) (*EndTaskResponse, error) {
	out := new(EndTaskResponse)
	err := c.cc.Invoke(ctx, "/LoggingService/EndTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoggingServiceServer is the server API for LoggingService service.
// All implementations must embed UnimplementedLoggingServiceServer
// for forward compatibility
type LoggingServiceServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	AppendLog(context.Context, *AppendLogRequest) (*AppendLogResponse, error)
	EndTask(context.Context, *EndTaskRequest) (*EndTaskResponse, error)
	mustEmbedUnimplementedLoggingServiceServer()
}

// UnimplementedLoggingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLoggingServiceServer struct {
}

func (UnimplementedLoggingServiceServer) CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedLoggingServiceServer) AppendLog(context.Context, *AppendLogRequest) (*AppendLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendLog not implemented")
}
func (UnimplementedLoggingServiceServer) EndTask(context.Context, *EndTaskRequest) (*EndTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndTask not implemented")
}
func (UnimplementedLoggingServiceServer) mustEmbedUnimplementedLoggingServiceServer() {}

// UnsafeLoggingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoggingServiceServer will
// result in compilation errors.
type UnsafeLoggingServiceServer interface {
	mustEmbedUnimplementedLoggingServiceServer()
}

func RegisterLoggingServiceServer(s grpc.ServiceRegistrar, srv LoggingServiceServer) {
	s.RegisterService(&LoggingService_ServiceDesc, srv)
}

func _LoggingService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LoggingService/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingService_AppendLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceServer).AppendLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LoggingService/AppendLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceServer).AppendLog(ctx, req.(*AppendLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingService_EndTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceServer).EndTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LoggingService/EndTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceServer).EndTask(ctx, req.(*EndTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoggingService_ServiceDesc is the grpc.ServiceDesc for LoggingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoggingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LoggingService",
	HandlerType: (*LoggingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _LoggingService_CreateTask_Handler,
		},
		{
			MethodName: "AppendLog",
			Handler:    _LoggingService_AppendLog_Handler,
		},
		{
			MethodName: "EndTask",
			Handler:    _LoggingService_EndTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/main.proto",
}
