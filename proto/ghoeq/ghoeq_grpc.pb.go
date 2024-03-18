// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: proto/ghoeq/ghoeq.proto

package ghoeq

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
	BackendServer_ListSources_FullMethodName        = "/ghoeq.BackendServer/ListSources"
	BackendServer_ListSession_FullMethodName        = "/ghoeq.BackendServer/ListSession"
	BackendServer_ListStreams_FullMethodName        = "/ghoeq.BackendServer/ListStreams"
	BackendServer_ListClients_FullMethodName        = "/ghoeq.BackendServer/ListClients"
	BackendServer_ModifySession_FullMethodName      = "/ghoeq.BackendServer/ModifySession"
	BackendServer_AttachSessionRaw_FullMethodName   = "/ghoeq.BackendServer/AttachSessionRaw"
	BackendServer_AttachClientStream_FullMethodName = "/ghoeq.BackendServer/AttachClientStream"
	BackendServer_AttachClient_FullMethodName       = "/ghoeq.BackendServer/AttachClient"
	BackendServer_AttachStreamRaw_FullMethodName    = "/ghoeq.BackendServer/AttachStreamRaw"
	BackendServer_AttachStreamStruct_FullMethodName = "/ghoeq.BackendServer/AttachStreamStruct"
)

// BackendServerClient is the client API for BackendServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BackendServerClient interface {
	ListSources(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListSourcesResponse, error)
	ListSession(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListSessionResponse, error)
	ListStreams(ctx context.Context, in *ListStreamRequest, opts ...grpc.CallOption) (*ListStreamsResponse, error)
	ListClients(ctx context.Context, in *ListClientRequest, opts ...grpc.CallOption) (*ListClientsResponse, error)
	ModifySession(ctx context.Context, in *ModifySessionRequest, opts ...grpc.CallOption) (*ModifySessionResponse, error)
	AttachSessionRaw(ctx context.Context, in *AttachSessionRequest, opts ...grpc.CallOption) (BackendServer_AttachSessionRawClient, error)
	AttachClientStream(ctx context.Context, in *AttachClientStreamRequest, opts ...grpc.CallOption) (BackendServer_AttachClientStreamClient, error)
	AttachClient(ctx context.Context, in *AttachClientRequest, opts ...grpc.CallOption) (BackendServer_AttachClientClient, error)
	AttachStreamRaw(ctx context.Context, in *AttachStreamRequest, opts ...grpc.CallOption) (BackendServer_AttachStreamRawClient, error)
	AttachStreamStruct(ctx context.Context, in *AttachStreamRequest, opts ...grpc.CallOption) (BackendServer_AttachStreamStructClient, error)
}

type backendServerClient struct {
	cc grpc.ClientConnInterface
}

func NewBackendServerClient(cc grpc.ClientConnInterface) BackendServerClient {
	return &backendServerClient{cc}
}

func (c *backendServerClient) ListSources(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListSourcesResponse, error) {
	out := new(ListSourcesResponse)
	err := c.cc.Invoke(ctx, BackendServer_ListSources_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendServerClient) ListSession(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListSessionResponse, error) {
	out := new(ListSessionResponse)
	err := c.cc.Invoke(ctx, BackendServer_ListSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendServerClient) ListStreams(ctx context.Context, in *ListStreamRequest, opts ...grpc.CallOption) (*ListStreamsResponse, error) {
	out := new(ListStreamsResponse)
	err := c.cc.Invoke(ctx, BackendServer_ListStreams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendServerClient) ListClients(ctx context.Context, in *ListClientRequest, opts ...grpc.CallOption) (*ListClientsResponse, error) {
	out := new(ListClientsResponse)
	err := c.cc.Invoke(ctx, BackendServer_ListClients_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendServerClient) ModifySession(ctx context.Context, in *ModifySessionRequest, opts ...grpc.CallOption) (*ModifySessionResponse, error) {
	out := new(ModifySessionResponse)
	err := c.cc.Invoke(ctx, BackendServer_ModifySession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendServerClient) AttachSessionRaw(ctx context.Context, in *AttachSessionRequest, opts ...grpc.CallOption) (BackendServer_AttachSessionRawClient, error) {
	stream, err := c.cc.NewStream(ctx, &BackendServer_ServiceDesc.Streams[0], BackendServer_AttachSessionRaw_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &backendServerAttachSessionRawClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BackendServer_AttachSessionRawClient interface {
	Recv() (*APPacket, error)
	grpc.ClientStream
}

type backendServerAttachSessionRawClient struct {
	grpc.ClientStream
}

func (x *backendServerAttachSessionRawClient) Recv() (*APPacket, error) {
	m := new(APPacket)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *backendServerClient) AttachClientStream(ctx context.Context, in *AttachClientStreamRequest, opts ...grpc.CallOption) (BackendServer_AttachClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &BackendServer_ServiceDesc.Streams[1], BackendServer_AttachClientStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &backendServerAttachClientStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BackendServer_AttachClientStreamClient interface {
	Recv() (*ClientPacket, error)
	grpc.ClientStream
}

type backendServerAttachClientStreamClient struct {
	grpc.ClientStream
}

func (x *backendServerAttachClientStreamClient) Recv() (*ClientPacket, error) {
	m := new(ClientPacket)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *backendServerClient) AttachClient(ctx context.Context, in *AttachClientRequest, opts ...grpc.CallOption) (BackendServer_AttachClientClient, error) {
	stream, err := c.cc.NewStream(ctx, &BackendServer_ServiceDesc.Streams[2], BackendServer_AttachClient_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &backendServerAttachClientClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BackendServer_AttachClientClient interface {
	Recv() (*ClientUpdate, error)
	grpc.ClientStream
}

type backendServerAttachClientClient struct {
	grpc.ClientStream
}

func (x *backendServerAttachClientClient) Recv() (*ClientUpdate, error) {
	m := new(ClientUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *backendServerClient) AttachStreamRaw(ctx context.Context, in *AttachStreamRequest, opts ...grpc.CallOption) (BackendServer_AttachStreamRawClient, error) {
	stream, err := c.cc.NewStream(ctx, &BackendServer_ServiceDesc.Streams[3], BackendServer_AttachStreamRaw_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &backendServerAttachStreamRawClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BackendServer_AttachStreamRawClient interface {
	Recv() (*APPacket, error)
	grpc.ClientStream
}

type backendServerAttachStreamRawClient struct {
	grpc.ClientStream
}

func (x *backendServerAttachStreamRawClient) Recv() (*APPacket, error) {
	m := new(APPacket)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *backendServerClient) AttachStreamStruct(ctx context.Context, in *AttachStreamRequest, opts ...grpc.CallOption) (BackendServer_AttachStreamStructClient, error) {
	stream, err := c.cc.NewStream(ctx, &BackendServer_ServiceDesc.Streams[4], BackendServer_AttachStreamStruct_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &backendServerAttachStreamStructClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BackendServer_AttachStreamStructClient interface {
	Recv() (*ClientPacket, error)
	grpc.ClientStream
}

type backendServerAttachStreamStructClient struct {
	grpc.ClientStream
}

func (x *backendServerAttachStreamStructClient) Recv() (*ClientPacket, error) {
	m := new(ClientPacket)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BackendServerServer is the server API for BackendServer service.
// All implementations must embed UnimplementedBackendServerServer
// for forward compatibility
type BackendServerServer interface {
	ListSources(context.Context, *ListRequest) (*ListSourcesResponse, error)
	ListSession(context.Context, *ListRequest) (*ListSessionResponse, error)
	ListStreams(context.Context, *ListStreamRequest) (*ListStreamsResponse, error)
	ListClients(context.Context, *ListClientRequest) (*ListClientsResponse, error)
	ModifySession(context.Context, *ModifySessionRequest) (*ModifySessionResponse, error)
	AttachSessionRaw(*AttachSessionRequest, BackendServer_AttachSessionRawServer) error
	AttachClientStream(*AttachClientStreamRequest, BackendServer_AttachClientStreamServer) error
	AttachClient(*AttachClientRequest, BackendServer_AttachClientServer) error
	AttachStreamRaw(*AttachStreamRequest, BackendServer_AttachStreamRawServer) error
	AttachStreamStruct(*AttachStreamRequest, BackendServer_AttachStreamStructServer) error
	mustEmbedUnimplementedBackendServerServer()
}

// UnimplementedBackendServerServer must be embedded to have forward compatible implementations.
type UnimplementedBackendServerServer struct {
}

func (UnimplementedBackendServerServer) ListSources(context.Context, *ListRequest) (*ListSourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSources not implemented")
}
func (UnimplementedBackendServerServer) ListSession(context.Context, *ListRequest) (*ListSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSession not implemented")
}
func (UnimplementedBackendServerServer) ListStreams(context.Context, *ListStreamRequest) (*ListStreamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStreams not implemented")
}
func (UnimplementedBackendServerServer) ListClients(context.Context, *ListClientRequest) (*ListClientsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClients not implemented")
}
func (UnimplementedBackendServerServer) ModifySession(context.Context, *ModifySessionRequest) (*ModifySessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySession not implemented")
}
func (UnimplementedBackendServerServer) AttachSessionRaw(*AttachSessionRequest, BackendServer_AttachSessionRawServer) error {
	return status.Errorf(codes.Unimplemented, "method AttachSessionRaw not implemented")
}
func (UnimplementedBackendServerServer) AttachClientStream(*AttachClientStreamRequest, BackendServer_AttachClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AttachClientStream not implemented")
}
func (UnimplementedBackendServerServer) AttachClient(*AttachClientRequest, BackendServer_AttachClientServer) error {
	return status.Errorf(codes.Unimplemented, "method AttachClient not implemented")
}
func (UnimplementedBackendServerServer) AttachStreamRaw(*AttachStreamRequest, BackendServer_AttachStreamRawServer) error {
	return status.Errorf(codes.Unimplemented, "method AttachStreamRaw not implemented")
}
func (UnimplementedBackendServerServer) AttachStreamStruct(*AttachStreamRequest, BackendServer_AttachStreamStructServer) error {
	return status.Errorf(codes.Unimplemented, "method AttachStreamStruct not implemented")
}
func (UnimplementedBackendServerServer) mustEmbedUnimplementedBackendServerServer() {}

// UnsafeBackendServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackendServerServer will
// result in compilation errors.
type UnsafeBackendServerServer interface {
	mustEmbedUnimplementedBackendServerServer()
}

func RegisterBackendServerServer(s grpc.ServiceRegistrar, srv BackendServerServer) {
	s.RegisterService(&BackendServer_ServiceDesc, srv)
}

func _BackendServer_ListSources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServerServer).ListSources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackendServer_ListSources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServerServer).ListSources(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackendServer_ListSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServerServer).ListSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackendServer_ListSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServerServer).ListSession(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackendServer_ListStreams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServerServer).ListStreams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackendServer_ListStreams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServerServer).ListStreams(ctx, req.(*ListStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackendServer_ListClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServerServer).ListClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackendServer_ListClients_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServerServer).ListClients(ctx, req.(*ListClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackendServer_ModifySession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServerServer).ModifySession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackendServer_ModifySession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServerServer).ModifySession(ctx, req.(*ModifySessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackendServer_AttachSessionRaw_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AttachSessionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BackendServerServer).AttachSessionRaw(m, &backendServerAttachSessionRawServer{stream})
}

type BackendServer_AttachSessionRawServer interface {
	Send(*APPacket) error
	grpc.ServerStream
}

type backendServerAttachSessionRawServer struct {
	grpc.ServerStream
}

func (x *backendServerAttachSessionRawServer) Send(m *APPacket) error {
	return x.ServerStream.SendMsg(m)
}

func _BackendServer_AttachClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AttachClientStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BackendServerServer).AttachClientStream(m, &backendServerAttachClientStreamServer{stream})
}

type BackendServer_AttachClientStreamServer interface {
	Send(*ClientPacket) error
	grpc.ServerStream
}

type backendServerAttachClientStreamServer struct {
	grpc.ServerStream
}

func (x *backendServerAttachClientStreamServer) Send(m *ClientPacket) error {
	return x.ServerStream.SendMsg(m)
}

func _BackendServer_AttachClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AttachClientRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BackendServerServer).AttachClient(m, &backendServerAttachClientServer{stream})
}

type BackendServer_AttachClientServer interface {
	Send(*ClientUpdate) error
	grpc.ServerStream
}

type backendServerAttachClientServer struct {
	grpc.ServerStream
}

func (x *backendServerAttachClientServer) Send(m *ClientUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func _BackendServer_AttachStreamRaw_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AttachStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BackendServerServer).AttachStreamRaw(m, &backendServerAttachStreamRawServer{stream})
}

type BackendServer_AttachStreamRawServer interface {
	Send(*APPacket) error
	grpc.ServerStream
}

type backendServerAttachStreamRawServer struct {
	grpc.ServerStream
}

func (x *backendServerAttachStreamRawServer) Send(m *APPacket) error {
	return x.ServerStream.SendMsg(m)
}

func _BackendServer_AttachStreamStruct_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AttachStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BackendServerServer).AttachStreamStruct(m, &backendServerAttachStreamStructServer{stream})
}

type BackendServer_AttachStreamStructServer interface {
	Send(*ClientPacket) error
	grpc.ServerStream
}

type backendServerAttachStreamStructServer struct {
	grpc.ServerStream
}

func (x *backendServerAttachStreamStructServer) Send(m *ClientPacket) error {
	return x.ServerStream.SendMsg(m)
}

// BackendServer_ServiceDesc is the grpc.ServiceDesc for BackendServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BackendServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ghoeq.BackendServer",
	HandlerType: (*BackendServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSources",
			Handler:    _BackendServer_ListSources_Handler,
		},
		{
			MethodName: "ListSession",
			Handler:    _BackendServer_ListSession_Handler,
		},
		{
			MethodName: "ListStreams",
			Handler:    _BackendServer_ListStreams_Handler,
		},
		{
			MethodName: "ListClients",
			Handler:    _BackendServer_ListClients_Handler,
		},
		{
			MethodName: "ModifySession",
			Handler:    _BackendServer_ModifySession_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AttachSessionRaw",
			Handler:       _BackendServer_AttachSessionRaw_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AttachClientStream",
			Handler:       _BackendServer_AttachClientStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AttachClient",
			Handler:       _BackendServer_AttachClient_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AttachStreamRaw",
			Handler:       _BackendServer_AttachStreamRaw_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AttachStreamStruct",
			Handler:       _BackendServer_AttachStreamStruct_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/ghoeq/ghoeq.proto",
}
