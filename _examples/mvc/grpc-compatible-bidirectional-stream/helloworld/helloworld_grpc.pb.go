// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package helloworld

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}

// GreeterServerSideSStreamClient is the client API for GreeterServerSideSStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterServerSideSStreamClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (GreeterServerSideSStream_SayHelloClient, error)
}

type greeterServerSideSStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterServerSideSStreamClient(cc grpc.ClientConnInterface) GreeterServerSideSStreamClient {
	return &greeterServerSideSStreamClient{cc}
}

func (c *greeterServerSideSStreamClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (GreeterServerSideSStream_SayHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreeterServerSideSStream_serviceDesc.Streams[0], "/helloworld.GreeterServerSideSStream/SayHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterServerSideSStreamSayHelloClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreeterServerSideSStream_SayHelloClient interface {
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterServerSideSStreamSayHelloClient struct {
	grpc.ClientStream
}

func (x *greeterServerSideSStreamSayHelloClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServerSideSStreamServer is the server API for GreeterServerSideSStream service.
// All implementations must embed UnimplementedGreeterServerSideSStreamServer
// for forward compatibility
type GreeterServerSideSStreamServer interface {
	// Sends a greeting
	SayHello(*HelloRequest, GreeterServerSideSStream_SayHelloServer) error
	mustEmbedUnimplementedGreeterServerSideSStreamServer()
}

// UnimplementedGreeterServerSideSStreamServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServerSideSStreamServer struct {
}

func (UnimplementedGreeterServerSideSStreamServer) SayHello(*HelloRequest, GreeterServerSideSStream_SayHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServerSideSStreamServer) mustEmbedUnimplementedGreeterServerSideSStreamServer() {
}

// UnsafeGreeterServerSideSStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServerSideSStreamServer will
// result in compilation errors.
type UnsafeGreeterServerSideSStreamServer interface {
	mustEmbedUnimplementedGreeterServerSideSStreamServer()
}

func RegisterGreeterServerSideSStreamServer(s grpc.ServiceRegistrar, srv GreeterServerSideSStreamServer) {
	s.RegisterService(&_GreeterServerSideSStream_serviceDesc, srv)
}

func _GreeterServerSideSStream_SayHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServerSideSStreamServer).SayHello(m, &greeterServerSideSStreamSayHelloServer{stream})
}

type GreeterServerSideSStream_SayHelloServer interface {
	Send(*HelloReply) error
	grpc.ServerStream
}

type greeterServerSideSStreamSayHelloServer struct {
	grpc.ServerStream
}

func (x *greeterServerSideSStreamSayHelloServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

var _GreeterServerSideSStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.GreeterServerSideSStream",
	HandlerType: (*GreeterServerSideSStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHello",
			Handler:       _GreeterServerSideSStream_SayHello_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}

// GreeterClientSideStreamClient is the client API for GreeterClientSideStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClientSideStreamClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, opts ...grpc.CallOption) (GreeterClientSideStream_SayHelloClient, error)
}

type greeterClientSideStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClientSideStreamClient(cc grpc.ClientConnInterface) GreeterClientSideStreamClient {
	return &greeterClientSideStreamClient{cc}
}

func (c *greeterClientSideStreamClient) SayHello(ctx context.Context, opts ...grpc.CallOption) (GreeterClientSideStream_SayHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreeterClientSideStream_serviceDesc.Streams[0], "/helloworld.GreeterClientSideStream/SayHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterClientSideStreamSayHelloClient{stream}
	return x, nil
}

type GreeterClientSideStream_SayHelloClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterClientSideStreamSayHelloClient struct {
	grpc.ClientStream
}

func (x *greeterClientSideStreamSayHelloClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterClientSideStreamSayHelloClient) CloseAndRecv() (*HelloReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterClientSideStreamServer is the server API for GreeterClientSideStream service.
// All implementations must embed UnimplementedGreeterClientSideStreamServer
// for forward compatibility
type GreeterClientSideStreamServer interface {
	// Sends a greeting
	SayHello(GreeterClientSideStream_SayHelloServer) error
	mustEmbedUnimplementedGreeterClientSideStreamServer()
}

// UnimplementedGreeterClientSideStreamServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterClientSideStreamServer struct {
}

func (UnimplementedGreeterClientSideStreamServer) SayHello(GreeterClientSideStream_SayHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterClientSideStreamServer) mustEmbedUnimplementedGreeterClientSideStreamServer() {
}

// UnsafeGreeterClientSideStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterClientSideStreamServer will
// result in compilation errors.
type UnsafeGreeterClientSideStreamServer interface {
	mustEmbedUnimplementedGreeterClientSideStreamServer()
}

func RegisterGreeterClientSideStreamServer(s grpc.ServiceRegistrar, srv GreeterClientSideStreamServer) {
	s.RegisterService(&_GreeterClientSideStream_serviceDesc, srv)
}

func _GreeterClientSideStream_SayHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterClientSideStreamServer).SayHello(&greeterClientSideStreamSayHelloServer{stream})
}

type GreeterClientSideStream_SayHelloServer interface {
	SendAndClose(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterClientSideStreamSayHelloServer struct {
	grpc.ServerStream
}

func (x *greeterClientSideStreamSayHelloServer) SendAndClose(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterClientSideStreamSayHelloServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GreeterClientSideStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.GreeterClientSideStream",
	HandlerType: (*GreeterClientSideStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHello",
			Handler:       _GreeterClientSideStream_SayHello_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}

// GreeterBidirectionalStreamClient is the client API for GreeterBidirectionalStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterBidirectionalStreamClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, opts ...grpc.CallOption) (GreeterBidirectionalStream_SayHelloClient, error)
}

type greeterBidirectionalStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterBidirectionalStreamClient(cc grpc.ClientConnInterface) GreeterBidirectionalStreamClient {
	return &greeterBidirectionalStreamClient{cc}
}

func (c *greeterBidirectionalStreamClient) SayHello(ctx context.Context, opts ...grpc.CallOption) (GreeterBidirectionalStream_SayHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreeterBidirectionalStream_serviceDesc.Streams[0], "/helloworld.GreeterBidirectionalStream/SayHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterBidirectionalStreamSayHelloClient{stream}
	return x, nil
}

type GreeterBidirectionalStream_SayHelloClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterBidirectionalStreamSayHelloClient struct {
	grpc.ClientStream
}

func (x *greeterBidirectionalStreamSayHelloClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterBidirectionalStreamSayHelloClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterBidirectionalStreamServer is the server API for GreeterBidirectionalStream service.
// All implementations must embed UnimplementedGreeterBidirectionalStreamServer
// for forward compatibility
type GreeterBidirectionalStreamServer interface {
	// Sends a greeting
	SayHello(GreeterBidirectionalStream_SayHelloServer) error
	mustEmbedUnimplementedGreeterBidirectionalStreamServer()
}

// UnimplementedGreeterBidirectionalStreamServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterBidirectionalStreamServer struct {
}

func (UnimplementedGreeterBidirectionalStreamServer) SayHello(GreeterBidirectionalStream_SayHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterBidirectionalStreamServer) mustEmbedUnimplementedGreeterBidirectionalStreamServer() {
}

// UnsafeGreeterBidirectionalStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterBidirectionalStreamServer will
// result in compilation errors.
type UnsafeGreeterBidirectionalStreamServer interface {
	mustEmbedUnimplementedGreeterBidirectionalStreamServer()
}

func RegisterGreeterBidirectionalStreamServer(s grpc.ServiceRegistrar, srv GreeterBidirectionalStreamServer) {
	s.RegisterService(&_GreeterBidirectionalStream_serviceDesc, srv)
}

func _GreeterBidirectionalStream_SayHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterBidirectionalStreamServer).SayHello(&greeterBidirectionalStreamSayHelloServer{stream})
}

type GreeterBidirectionalStream_SayHelloServer interface {
	Send(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterBidirectionalStreamSayHelloServer struct {
	grpc.ServerStream
}

func (x *greeterBidirectionalStreamSayHelloServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterBidirectionalStreamSayHelloServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GreeterBidirectionalStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.GreeterBidirectionalStream",
	HandlerType: (*GreeterBidirectionalStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHello",
			Handler:       _GreeterBidirectionalStream_SayHello_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}
