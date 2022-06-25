// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pipes

import (
	context "context"
	messages "github.com/nochte/pipelinr-protocol/protobuf/messages"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PipeClient is the client API for Pipe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PipeClient interface {
	StreamSend(ctx context.Context, opts ...grpc.CallOption) (Pipe_StreamSendClient, error)
	// Send will send up a message envelop, and will return an event id, error if invalid for any reason
	Send(ctx context.Context, in *messages.MessageEnvelop, opts ...grpc.CallOption) (*Xid, error)
	// Recv will request to receive with options. Defaults to {
	//     AutoAck = false,
	//     Block = true,
	//     Count = 1,
	//     Timeout = inf
	// }
	Recv(ctx context.Context, in *ReceiveOptions, opts ...grpc.CallOption) (*messages.Events, error)
	StreamRecv(ctx context.Context, in *ReceiveOptions, opts ...grpc.CallOption) (Pipe_StreamRecvClient, error)
	// Ack acknowledges that a message by id was received and can be discarded from the re-enqueue queue queue
	Ack(ctx context.Context, in *CompleteRequest, opts ...grpc.CallOption) (*GenericResponse, error)
	// Complete takes a Xid and step, marking the step as complete (to be enqueued into the next pipe, if needed)
	Complete(ctx context.Context, in *CompleteRequest, opts ...grpc.CallOption) (*GenericResponse, error)
	// AppendLog takes a routelog and adds it to the message. If step is not given, assumes current step
	AppendLog(ctx context.Context, in *RouteLogRequest, opts ...grpc.CallOption) (*GenericResponse, error)
	// AddSteps adds steps to the route. If After is not given, assumes after current step
	AddSteps(ctx context.Context, in *AddStepsRequest, opts ...grpc.CallOption) (*GenericResponse, error)
	// Decorate takes a set of decorations and applies them to the message
	Decorate(ctx context.Context, in *Decorations, opts ...grpc.CallOption) (*GenericResponses, error)
	// GetDecoration, given a message id and decoration keys, yields the values of those field
	GetDecoration(ctx context.Context, in *GetDecorationRequest, opts ...grpc.CallOption) (*Decoration, error)
}

type pipeClient struct {
	cc grpc.ClientConnInterface
}

func NewPipeClient(cc grpc.ClientConnInterface) PipeClient {
	return &pipeClient{cc}
}

func (c *pipeClient) StreamSend(ctx context.Context, opts ...grpc.CallOption) (Pipe_StreamSendClient, error) {
	stream, err := c.cc.NewStream(ctx, &Pipe_ServiceDesc.Streams[0], "/pipes.Pipe/StreamSend", opts...)
	if err != nil {
		return nil, err
	}
	x := &pipeStreamSendClient{stream}
	return x, nil
}

type Pipe_StreamSendClient interface {
	Send(*messages.MessageEnvelop) error
	Recv() (*Xid, error)
	grpc.ClientStream
}

type pipeStreamSendClient struct {
	grpc.ClientStream
}

func (x *pipeStreamSendClient) Send(m *messages.MessageEnvelop) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pipeStreamSendClient) Recv() (*Xid, error) {
	m := new(Xid)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pipeClient) Send(ctx context.Context, in *messages.MessageEnvelop, opts ...grpc.CallOption) (*Xid, error) {
	out := new(Xid)
	err := c.cc.Invoke(ctx, "/pipes.Pipe/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipeClient) Recv(ctx context.Context, in *ReceiveOptions, opts ...grpc.CallOption) (*messages.Events, error) {
	out := new(messages.Events)
	err := c.cc.Invoke(ctx, "/pipes.Pipe/Recv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipeClient) StreamRecv(ctx context.Context, in *ReceiveOptions, opts ...grpc.CallOption) (Pipe_StreamRecvClient, error) {
	stream, err := c.cc.NewStream(ctx, &Pipe_ServiceDesc.Streams[1], "/pipes.Pipe/StreamRecv", opts...)
	if err != nil {
		return nil, err
	}
	x := &pipeStreamRecvClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Pipe_StreamRecvClient interface {
	Recv() (*messages.Event, error)
	grpc.ClientStream
}

type pipeStreamRecvClient struct {
	grpc.ClientStream
}

func (x *pipeStreamRecvClient) Recv() (*messages.Event, error) {
	m := new(messages.Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pipeClient) Ack(ctx context.Context, in *CompleteRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/pipes.Pipe/Ack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipeClient) Complete(ctx context.Context, in *CompleteRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/pipes.Pipe/Complete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipeClient) AppendLog(ctx context.Context, in *RouteLogRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/pipes.Pipe/AppendLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipeClient) AddSteps(ctx context.Context, in *AddStepsRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/pipes.Pipe/AddSteps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipeClient) Decorate(ctx context.Context, in *Decorations, opts ...grpc.CallOption) (*GenericResponses, error) {
	out := new(GenericResponses)
	err := c.cc.Invoke(ctx, "/pipes.Pipe/Decorate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipeClient) GetDecoration(ctx context.Context, in *GetDecorationRequest, opts ...grpc.CallOption) (*Decoration, error) {
	out := new(Decoration)
	err := c.cc.Invoke(ctx, "/pipes.Pipe/GetDecoration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PipeServer is the server API for Pipe service.
// All implementations should embed UnimplementedPipeServer
// for forward compatibility
type PipeServer interface {
	StreamSend(Pipe_StreamSendServer) error
	// Send will send up a message envelop, and will return an event id, error if invalid for any reason
	Send(context.Context, *messages.MessageEnvelop) (*Xid, error)
	// Recv will request to receive with options. Defaults to {
	//     AutoAck = false,
	//     Block = true,
	//     Count = 1,
	//     Timeout = inf
	// }
	Recv(context.Context, *ReceiveOptions) (*messages.Events, error)
	StreamRecv(*ReceiveOptions, Pipe_StreamRecvServer) error
	// Ack acknowledges that a message by id was received and can be discarded from the re-enqueue queue queue
	Ack(context.Context, *CompleteRequest) (*GenericResponse, error)
	// Complete takes a Xid and step, marking the step as complete (to be enqueued into the next pipe, if needed)
	Complete(context.Context, *CompleteRequest) (*GenericResponse, error)
	// AppendLog takes a routelog and adds it to the message. If step is not given, assumes current step
	AppendLog(context.Context, *RouteLogRequest) (*GenericResponse, error)
	// AddSteps adds steps to the route. If After is not given, assumes after current step
	AddSteps(context.Context, *AddStepsRequest) (*GenericResponse, error)
	// Decorate takes a set of decorations and applies them to the message
	Decorate(context.Context, *Decorations) (*GenericResponses, error)
	// GetDecoration, given a message id and decoration keys, yields the values of those field
	GetDecoration(context.Context, *GetDecorationRequest) (*Decoration, error)
}

// UnimplementedPipeServer should be embedded to have forward compatible implementations.
type UnimplementedPipeServer struct {
}

func (UnimplementedPipeServer) StreamSend(Pipe_StreamSendServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamSend not implemented")
}
func (UnimplementedPipeServer) Send(context.Context, *messages.MessageEnvelop) (*Xid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedPipeServer) Recv(context.Context, *ReceiveOptions) (*messages.Events, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recv not implemented")
}
func (UnimplementedPipeServer) StreamRecv(*ReceiveOptions, Pipe_StreamRecvServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRecv not implemented")
}
func (UnimplementedPipeServer) Ack(context.Context, *CompleteRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ack not implemented")
}
func (UnimplementedPipeServer) Complete(context.Context, *CompleteRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Complete not implemented")
}
func (UnimplementedPipeServer) AppendLog(context.Context, *RouteLogRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendLog not implemented")
}
func (UnimplementedPipeServer) AddSteps(context.Context, *AddStepsRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSteps not implemented")
}
func (UnimplementedPipeServer) Decorate(context.Context, *Decorations) (*GenericResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decorate not implemented")
}
func (UnimplementedPipeServer) GetDecoration(context.Context, *GetDecorationRequest) (*Decoration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDecoration not implemented")
}

// UnsafePipeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PipeServer will
// result in compilation errors.
type UnsafePipeServer interface {
	mustEmbedUnimplementedPipeServer()
}

func RegisterPipeServer(s grpc.ServiceRegistrar, srv PipeServer) {
	s.RegisterService(&Pipe_ServiceDesc, srv)
}

func _Pipe_StreamSend_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PipeServer).StreamSend(&pipeStreamSendServer{stream})
}

type Pipe_StreamSendServer interface {
	Send(*Xid) error
	Recv() (*messages.MessageEnvelop, error)
	grpc.ServerStream
}

type pipeStreamSendServer struct {
	grpc.ServerStream
}

func (x *pipeStreamSendServer) Send(m *Xid) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pipeStreamSendServer) Recv() (*messages.MessageEnvelop, error) {
	m := new(messages.MessageEnvelop)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Pipe_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(messages.MessageEnvelop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipeServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pipes.Pipe/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipeServer).Send(ctx, req.(*messages.MessageEnvelop))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipe_Recv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipeServer).Recv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pipes.Pipe/Recv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipeServer).Recv(ctx, req.(*ReceiveOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipe_StreamRecv_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveOptions)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PipeServer).StreamRecv(m, &pipeStreamRecvServer{stream})
}

type Pipe_StreamRecvServer interface {
	Send(*messages.Event) error
	grpc.ServerStream
}

type pipeStreamRecvServer struct {
	grpc.ServerStream
}

func (x *pipeStreamRecvServer) Send(m *messages.Event) error {
	return x.ServerStream.SendMsg(m)
}

func _Pipe_Ack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipeServer).Ack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pipes.Pipe/Ack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipeServer).Ack(ctx, req.(*CompleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipe_Complete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipeServer).Complete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pipes.Pipe/Complete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipeServer).Complete(ctx, req.(*CompleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipe_AppendLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipeServer).AppendLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pipes.Pipe/AppendLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipeServer).AppendLog(ctx, req.(*RouteLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipe_AddSteps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStepsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipeServer).AddSteps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pipes.Pipe/AddSteps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipeServer).AddSteps(ctx, req.(*AddStepsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipe_Decorate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Decorations)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipeServer).Decorate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pipes.Pipe/Decorate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipeServer).Decorate(ctx, req.(*Decorations))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pipe_GetDecoration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDecorationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipeServer).GetDecoration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pipes.Pipe/GetDecoration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipeServer).GetDecoration(ctx, req.(*GetDecorationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Pipe_ServiceDesc is the grpc.ServiceDesc for Pipe service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pipe_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pipes.Pipe",
	HandlerType: (*PipeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Pipe_Send_Handler,
		},
		{
			MethodName: "Recv",
			Handler:    _Pipe_Recv_Handler,
		},
		{
			MethodName: "Ack",
			Handler:    _Pipe_Ack_Handler,
		},
		{
			MethodName: "Complete",
			Handler:    _Pipe_Complete_Handler,
		},
		{
			MethodName: "AppendLog",
			Handler:    _Pipe_AppendLog_Handler,
		},
		{
			MethodName: "AddSteps",
			Handler:    _Pipe_AddSteps_Handler,
		},
		{
			MethodName: "Decorate",
			Handler:    _Pipe_Decorate_Handler,
		},
		{
			MethodName: "GetDecoration",
			Handler:    _Pipe_GetDecoration_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamSend",
			Handler:       _Pipe_StreamSend_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamRecv",
			Handler:       _Pipe_StreamRecv_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pipe.proto",
}

// StoreClient is the client API for Store service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreClient interface {
	Get(ctx context.Context, in *Xid, opts ...grpc.CallOption) (*messages.Event, error)
	Del(ctx context.Context, in *Xid, opts ...grpc.CallOption) (*GenericResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*messages.Events, error)
}

type storeClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreClient(cc grpc.ClientConnInterface) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) Get(ctx context.Context, in *Xid, opts ...grpc.CallOption) (*messages.Event, error) {
	out := new(messages.Event)
	err := c.cc.Invoke(ctx, "/pipes.Store/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Del(ctx context.Context, in *Xid, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/pipes.Store/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*messages.Events, error) {
	out := new(messages.Events)
	err := c.cc.Invoke(ctx, "/pipes.Store/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServer is the server API for Store service.
// All implementations should embed UnimplementedStoreServer
// for forward compatibility
type StoreServer interface {
	Get(context.Context, *Xid) (*messages.Event, error)
	Del(context.Context, *Xid) (*GenericResponse, error)
	List(context.Context, *ListRequest) (*messages.Events, error)
}

// UnimplementedStoreServer should be embedded to have forward compatible implementations.
type UnimplementedStoreServer struct {
}

func (UnimplementedStoreServer) Get(context.Context, *Xid) (*messages.Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStoreServer) Del(context.Context, *Xid) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (UnimplementedStoreServer) List(context.Context, *ListRequest) (*messages.Events, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

// UnsafeStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreServer will
// result in compilation errors.
type UnsafeStoreServer interface {
	mustEmbedUnimplementedStoreServer()
}

func RegisterStoreServer(s grpc.ServiceRegistrar, srv StoreServer) {
	s.RegisterService(&Store_ServiceDesc, srv)
}

func _Store_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Xid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pipes.Store/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Get(ctx, req.(*Xid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Xid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pipes.Store/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Del(ctx, req.(*Xid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pipes.Store/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Store_ServiceDesc is the grpc.ServiceDesc for Store service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Store_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pipes.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Store_Get_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _Store_Del_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Store_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pipe.proto",
}
