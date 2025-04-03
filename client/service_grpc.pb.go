// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: service.proto

package client

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Command_Set_FullMethodName = "/client.Command/Set"
	Command_Get_FullMethodName = "/client.Command/Get"
	Command_Del_FullMethodName = "/client.Command/Del"
)

// CommandClient is the client API for Command service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandClient interface {
	Set(ctx context.Context, in *SetCmdRequest, opts ...grpc.CallOption) (*SetCmdResponse, error)
	Get(ctx context.Context, in *GetCmdRequest, opts ...grpc.CallOption) (*GetCmdResponse, error)
	Del(ctx context.Context, in *DeleteCmdRequest, opts ...grpc.CallOption) (*DeleteCmdResponse, error)
}

type commandClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandClient(cc grpc.ClientConnInterface) CommandClient {
	return &commandClient{cc}
}

func (c *commandClient) Set(ctx context.Context, in *SetCmdRequest, opts ...grpc.CallOption) (*SetCmdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetCmdResponse)
	err := c.cc.Invoke(ctx, Command_Set_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandClient) Get(ctx context.Context, in *GetCmdRequest, opts ...grpc.CallOption) (*GetCmdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCmdResponse)
	err := c.cc.Invoke(ctx, Command_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandClient) Del(ctx context.Context, in *DeleteCmdRequest, opts ...grpc.CallOption) (*DeleteCmdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCmdResponse)
	err := c.cc.Invoke(ctx, Command_Del_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandServer is the server API for Command service.
// All implementations must embed UnimplementedCommandServer
// for forward compatibility.
type CommandServer interface {
	Set(context.Context, *SetCmdRequest) (*SetCmdResponse, error)
	Get(context.Context, *GetCmdRequest) (*GetCmdResponse, error)
	Del(context.Context, *DeleteCmdRequest) (*DeleteCmdResponse, error)
	mustEmbedUnimplementedCommandServer()
}

// UnimplementedCommandServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCommandServer struct{}

func (UnimplementedCommandServer) Set(context.Context, *SetCmdRequest) (*SetCmdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedCommandServer) Get(context.Context, *GetCmdRequest) (*GetCmdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCommandServer) Del(context.Context, *DeleteCmdRequest) (*DeleteCmdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (UnimplementedCommandServer) mustEmbedUnimplementedCommandServer() {}
func (UnimplementedCommandServer) testEmbeddedByValue()                 {}

// UnsafeCommandServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommandServer will
// result in compilation errors.
type UnsafeCommandServer interface {
	mustEmbedUnimplementedCommandServer()
}

func RegisterCommandServer(s grpc.ServiceRegistrar, srv CommandServer) {
	// If the following call pancis, it indicates UnimplementedCommandServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Command_ServiceDesc, srv)
}

func _Command_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCmdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Command_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).Set(ctx, req.(*SetCmdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Command_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCmdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Command_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).Get(ctx, req.(*GetCmdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Command_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCmdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Command_Del_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).Del(ctx, req.(*DeleteCmdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Command_ServiceDesc is the grpc.ServiceDesc for Command service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Command_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "client.Command",
	HandlerType: (*CommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _Command_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Command_Get_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _Command_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
