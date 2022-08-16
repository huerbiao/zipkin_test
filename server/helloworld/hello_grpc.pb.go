// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: hello.proto

package helloworld

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

// HellowClient is the client API for Hellow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HellowClient interface {
	HelloWord(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error)
}

type hellowClient struct {
	cc grpc.ClientConnInterface
}

func NewHellowClient(cc grpc.ClientConnInterface) HellowClient {
	return &hellowClient{cc}
}

func (c *hellowClient) HelloWord(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error) {
	out := new(HelloResp)
	err := c.cc.Invoke(ctx, "/helloword.Hellow/HelloWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HellowServer is the server API for Hellow service.
// All implementations must embed UnimplementedHellowServer
// for forward compatibility
type HellowServer interface {
	HelloWord(context.Context, *HelloReq) (*HelloResp, error)
	mustEmbedUnimplementedHellowServer()
}

// UnimplementedHellowServer must be embedded to have forward compatible implementations.
type UnimplementedHellowServer struct {
}

func (UnimplementedHellowServer) HelloWord(context.Context, *HelloReq) (*HelloResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWord not implemented")
}
func (UnimplementedHellowServer) mustEmbedUnimplementedHellowServer() {}

// UnsafeHellowServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HellowServer will
// result in compilation errors.
type UnsafeHellowServer interface {
	mustEmbedUnimplementedHellowServer()
}

func RegisterHellowServer(s grpc.ServiceRegistrar, srv HellowServer) {
	s.RegisterService(&Hellow_ServiceDesc, srv)
}

func _Hellow_HelloWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HellowServer).HelloWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloword.Hellow/HelloWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HellowServer).HelloWord(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Hellow_ServiceDesc is the grpc.ServiceDesc for Hellow service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hellow_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloword.Hellow",
	HandlerType: (*HellowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWord",
			Handler:    _Hellow_HelloWord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}