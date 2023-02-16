// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.5.0
// source: push/push_server.proto

package push

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

// PushServerClient is the client API for PushServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PushServerClient interface {
	StartPush(ctx context.Context, in *StartPushReq, opts ...grpc.CallOption) (*StartPushRsp, error)
	// backend push
	BatchPush(ctx context.Context, in *BatchPushReq, opts ...grpc.CallOption) (*BatchPushRsp, error)
}

type pushServerClient struct {
	cc grpc.ClientConnInterface
}

func NewPushServerClient(cc grpc.ClientConnInterface) PushServerClient {
	return &pushServerClient{cc}
}

func (c *pushServerClient) StartPush(ctx context.Context, in *StartPushReq, opts ...grpc.CallOption) (*StartPushRsp, error) {
	out := new(StartPushRsp)
	err := c.cc.Invoke(ctx, "/push.PushServer/StartPush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushServerClient) BatchPush(ctx context.Context, in *BatchPushReq, opts ...grpc.CallOption) (*BatchPushRsp, error) {
	out := new(BatchPushRsp)
	err := c.cc.Invoke(ctx, "/push.PushServer/BatchPush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushServerServer is the server API for PushServer service.
// All implementations must embed UnimplementedPushServerServer
// for forward compatibility
type PushServerServer interface {
	StartPush(context.Context, *StartPushReq) (*StartPushRsp, error)
	// backend push
	BatchPush(context.Context, *BatchPushReq) (*BatchPushRsp, error)
	mustEmbedUnimplementedPushServerServer()
}

// UnimplementedPushServerServer must be embedded to have forward compatible implementations.
type UnimplementedPushServerServer struct {
}

func (UnimplementedPushServerServer) StartPush(context.Context, *StartPushReq) (*StartPushRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPush not implemented")
}
func (UnimplementedPushServerServer) BatchPush(context.Context, *BatchPushReq) (*BatchPushRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchPush not implemented")
}
func (UnimplementedPushServerServer) mustEmbedUnimplementedPushServerServer() {}

// UnsafePushServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PushServerServer will
// result in compilation errors.
type UnsafePushServerServer interface {
	mustEmbedUnimplementedPushServerServer()
}

func RegisterPushServerServer(s grpc.ServiceRegistrar, srv PushServerServer) {
	s.RegisterService(&PushServer_ServiceDesc, srv)
}

func _PushServer_StartPush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPushReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServerServer).StartPush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/push.PushServer/StartPush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServerServer).StartPush(ctx, req.(*StartPushReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PushServer_BatchPush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchPushReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServerServer).BatchPush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/push.PushServer/BatchPush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServerServer).BatchPush(ctx, req.(*BatchPushReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PushServer_ServiceDesc is the grpc.ServiceDesc for PushServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PushServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "push.PushServer",
	HandlerType: (*PushServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartPush",
			Handler:    _PushServer_StartPush_Handler,
		},
		{
			MethodName: "BatchPush",
			Handler:    _PushServer_BatchPush_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "push/push_server.proto",
}