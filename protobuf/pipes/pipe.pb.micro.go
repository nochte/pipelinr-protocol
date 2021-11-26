// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: pipe.proto

package pipes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	messages "github.com/nochte/pipelinr-protocol/protobuf/messages"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Pipe service

func NewPipeEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Pipe service

type PipeService interface {
	StreamSend(ctx context.Context, opts ...client.CallOption) (Pipe_StreamSendService, error)
	// Send will send up a message envelop, and will return an event id, error if invalid for any reason
	Send(ctx context.Context, in *messages.MessageEnvelop, opts ...client.CallOption) (*Xid, error)
	// Recv will request to receive with options. Defaults to {
	//     AutoAck = false,
	//     Block = true,
	//     Count = 1,
	//     Timeout = inf
	// }
	Recv(ctx context.Context, in *ReceiveOptions, opts ...client.CallOption) (*messages.Events, error)
	StreamRecv(ctx context.Context, in *ReceiveOptions, opts ...client.CallOption) (Pipe_StreamRecvService, error)
	// Ack acknowledges that a message by id was received and can be discarded from the re-enqueue queue queue
	Ack(ctx context.Context, in *CompleteRequest, opts ...client.CallOption) (*GenericResponse, error)
	// Complete takes a Xid and step, marking the step as complete (to be enqueued into the next pipe, if needed)
	Complete(ctx context.Context, in *CompleteRequest, opts ...client.CallOption) (*GenericResponse, error)
	// AppendLog takes a routelog and adds it to the message. If step is not given, assumes current step
	AppendLog(ctx context.Context, in *RouteLogRequest, opts ...client.CallOption) (*GenericResponse, error)
	// AddSteps adds steps to the route. If After is not given, assumes after current step
	AddSteps(ctx context.Context, in *AddStepsRequest, opts ...client.CallOption) (*GenericResponse, error)
	// Decorate takes a set of decorations and applies them to the message
	Decorate(ctx context.Context, in *Decorations, opts ...client.CallOption) (*GenericResponses, error)
}

type pipeService struct {
	c    client.Client
	name string
}

func NewPipeService(name string, c client.Client) PipeService {
	return &pipeService{
		c:    c,
		name: name,
	}
}

func (c *pipeService) StreamSend(ctx context.Context, opts ...client.CallOption) (Pipe_StreamSendService, error) {
	req := c.c.NewRequest(c.name, "Pipe.StreamSend", &messages.MessageEnvelop{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &pipeServiceStreamSend{stream}, nil
}

type Pipe_StreamSendService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*messages.MessageEnvelop) error
	Recv() (*Xid, error)
}

type pipeServiceStreamSend struct {
	stream client.Stream
}

func (x *pipeServiceStreamSend) Close() error {
	return x.stream.Close()
}

func (x *pipeServiceStreamSend) Context() context.Context {
	return x.stream.Context()
}

func (x *pipeServiceStreamSend) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *pipeServiceStreamSend) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *pipeServiceStreamSend) Send(m *messages.MessageEnvelop) error {
	return x.stream.Send(m)
}

func (x *pipeServiceStreamSend) Recv() (*Xid, error) {
	m := new(Xid)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pipeService) Send(ctx context.Context, in *messages.MessageEnvelop, opts ...client.CallOption) (*Xid, error) {
	req := c.c.NewRequest(c.name, "Pipe.Send", in)
	out := new(Xid)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipeService) Recv(ctx context.Context, in *ReceiveOptions, opts ...client.CallOption) (*messages.Events, error) {
	req := c.c.NewRequest(c.name, "Pipe.Recv", in)
	out := new(messages.Events)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipeService) StreamRecv(ctx context.Context, in *ReceiveOptions, opts ...client.CallOption) (Pipe_StreamRecvService, error) {
	req := c.c.NewRequest(c.name, "Pipe.StreamRecv", &ReceiveOptions{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &pipeServiceStreamRecv{stream}, nil
}

type Pipe_StreamRecvService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*messages.Event, error)
}

type pipeServiceStreamRecv struct {
	stream client.Stream
}

func (x *pipeServiceStreamRecv) Close() error {
	return x.stream.Close()
}

func (x *pipeServiceStreamRecv) Context() context.Context {
	return x.stream.Context()
}

func (x *pipeServiceStreamRecv) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *pipeServiceStreamRecv) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *pipeServiceStreamRecv) Recv() (*messages.Event, error) {
	m := new(messages.Event)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pipeService) Ack(ctx context.Context, in *CompleteRequest, opts ...client.CallOption) (*GenericResponse, error) {
	req := c.c.NewRequest(c.name, "Pipe.Ack", in)
	out := new(GenericResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipeService) Complete(ctx context.Context, in *CompleteRequest, opts ...client.CallOption) (*GenericResponse, error) {
	req := c.c.NewRequest(c.name, "Pipe.Complete", in)
	out := new(GenericResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipeService) AppendLog(ctx context.Context, in *RouteLogRequest, opts ...client.CallOption) (*GenericResponse, error) {
	req := c.c.NewRequest(c.name, "Pipe.AppendLog", in)
	out := new(GenericResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipeService) AddSteps(ctx context.Context, in *AddStepsRequest, opts ...client.CallOption) (*GenericResponse, error) {
	req := c.c.NewRequest(c.name, "Pipe.AddSteps", in)
	out := new(GenericResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipeService) Decorate(ctx context.Context, in *Decorations, opts ...client.CallOption) (*GenericResponses, error) {
	req := c.c.NewRequest(c.name, "Pipe.Decorate", in)
	out := new(GenericResponses)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Pipe service

type PipeHandler interface {
	StreamSend(context.Context, Pipe_StreamSendStream) error
	// Send will send up a message envelop, and will return an event id, error if invalid for any reason
	Send(context.Context, *messages.MessageEnvelop, *Xid) error
	// Recv will request to receive with options. Defaults to {
	//     AutoAck = false,
	//     Block = true,
	//     Count = 1,
	//     Timeout = inf
	// }
	Recv(context.Context, *ReceiveOptions, *messages.Events) error
	StreamRecv(context.Context, *ReceiveOptions, Pipe_StreamRecvStream) error
	// Ack acknowledges that a message by id was received and can be discarded from the re-enqueue queue queue
	Ack(context.Context, *CompleteRequest, *GenericResponse) error
	// Complete takes a Xid and step, marking the step as complete (to be enqueued into the next pipe, if needed)
	Complete(context.Context, *CompleteRequest, *GenericResponse) error
	// AppendLog takes a routelog and adds it to the message. If step is not given, assumes current step
	AppendLog(context.Context, *RouteLogRequest, *GenericResponse) error
	// AddSteps adds steps to the route. If After is not given, assumes after current step
	AddSteps(context.Context, *AddStepsRequest, *GenericResponse) error
	// Decorate takes a set of decorations and applies them to the message
	Decorate(context.Context, *Decorations, *GenericResponses) error
}

func RegisterPipeHandler(s server.Server, hdlr PipeHandler, opts ...server.HandlerOption) error {
	type pipe interface {
		StreamSend(ctx context.Context, stream server.Stream) error
		Send(ctx context.Context, in *messages.MessageEnvelop, out *Xid) error
		Recv(ctx context.Context, in *ReceiveOptions, out *messages.Events) error
		StreamRecv(ctx context.Context, stream server.Stream) error
		Ack(ctx context.Context, in *CompleteRequest, out *GenericResponse) error
		Complete(ctx context.Context, in *CompleteRequest, out *GenericResponse) error
		AppendLog(ctx context.Context, in *RouteLogRequest, out *GenericResponse) error
		AddSteps(ctx context.Context, in *AddStepsRequest, out *GenericResponse) error
		Decorate(ctx context.Context, in *Decorations, out *GenericResponses) error
	}
	type Pipe struct {
		pipe
	}
	h := &pipeHandler{hdlr}
	return s.Handle(s.NewHandler(&Pipe{h}, opts...))
}

type pipeHandler struct {
	PipeHandler
}

func (h *pipeHandler) StreamSend(ctx context.Context, stream server.Stream) error {
	return h.PipeHandler.StreamSend(ctx, &pipeStreamSendStream{stream})
}

type Pipe_StreamSendStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Xid) error
	Recv() (*messages.MessageEnvelop, error)
}

type pipeStreamSendStream struct {
	stream server.Stream
}

func (x *pipeStreamSendStream) Close() error {
	return x.stream.Close()
}

func (x *pipeStreamSendStream) Context() context.Context {
	return x.stream.Context()
}

func (x *pipeStreamSendStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *pipeStreamSendStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *pipeStreamSendStream) Send(m *Xid) error {
	return x.stream.Send(m)
}

func (x *pipeStreamSendStream) Recv() (*messages.MessageEnvelop, error) {
	m := new(messages.MessageEnvelop)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *pipeHandler) Send(ctx context.Context, in *messages.MessageEnvelop, out *Xid) error {
	return h.PipeHandler.Send(ctx, in, out)
}

func (h *pipeHandler) Recv(ctx context.Context, in *ReceiveOptions, out *messages.Events) error {
	return h.PipeHandler.Recv(ctx, in, out)
}

func (h *pipeHandler) StreamRecv(ctx context.Context, stream server.Stream) error {
	m := new(ReceiveOptions)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.PipeHandler.StreamRecv(ctx, m, &pipeStreamRecvStream{stream})
}

type Pipe_StreamRecvStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*messages.Event) error
}

type pipeStreamRecvStream struct {
	stream server.Stream
}

func (x *pipeStreamRecvStream) Close() error {
	return x.stream.Close()
}

func (x *pipeStreamRecvStream) Context() context.Context {
	return x.stream.Context()
}

func (x *pipeStreamRecvStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *pipeStreamRecvStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *pipeStreamRecvStream) Send(m *messages.Event) error {
	return x.stream.Send(m)
}

func (h *pipeHandler) Ack(ctx context.Context, in *CompleteRequest, out *GenericResponse) error {
	return h.PipeHandler.Ack(ctx, in, out)
}

func (h *pipeHandler) Complete(ctx context.Context, in *CompleteRequest, out *GenericResponse) error {
	return h.PipeHandler.Complete(ctx, in, out)
}

func (h *pipeHandler) AppendLog(ctx context.Context, in *RouteLogRequest, out *GenericResponse) error {
	return h.PipeHandler.AppendLog(ctx, in, out)
}

func (h *pipeHandler) AddSteps(ctx context.Context, in *AddStepsRequest, out *GenericResponse) error {
	return h.PipeHandler.AddSteps(ctx, in, out)
}

func (h *pipeHandler) Decorate(ctx context.Context, in *Decorations, out *GenericResponses) error {
	return h.PipeHandler.Decorate(ctx, in, out)
}

// Api Endpoints for Store service

func NewStoreEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Store service

type StoreService interface {
	Get(ctx context.Context, in *Xid, opts ...client.CallOption) (*messages.Event, error)
	Del(ctx context.Context, in *Xid, opts ...client.CallOption) (*GenericResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*messages.Events, error)
}

type storeService struct {
	c    client.Client
	name string
}

func NewStoreService(name string, c client.Client) StoreService {
	return &storeService{
		c:    c,
		name: name,
	}
}

func (c *storeService) Get(ctx context.Context, in *Xid, opts ...client.CallOption) (*messages.Event, error) {
	req := c.c.NewRequest(c.name, "Store.Get", in)
	out := new(messages.Event)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) Del(ctx context.Context, in *Xid, opts ...client.CallOption) (*GenericResponse, error) {
	req := c.c.NewRequest(c.name, "Store.Del", in)
	out := new(GenericResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*messages.Events, error) {
	req := c.c.NewRequest(c.name, "Store.List", in)
	out := new(messages.Events)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Store service

type StoreHandler interface {
	Get(context.Context, *Xid, *messages.Event) error
	Del(context.Context, *Xid, *GenericResponse) error
	List(context.Context, *ListRequest, *messages.Events) error
}

func RegisterStoreHandler(s server.Server, hdlr StoreHandler, opts ...server.HandlerOption) error {
	type store interface {
		Get(ctx context.Context, in *Xid, out *messages.Event) error
		Del(ctx context.Context, in *Xid, out *GenericResponse) error
		List(ctx context.Context, in *ListRequest, out *messages.Events) error
	}
	type Store struct {
		store
	}
	h := &storeHandler{hdlr}
	return s.Handle(s.NewHandler(&Store{h}, opts...))
}

type storeHandler struct {
	StoreHandler
}

func (h *storeHandler) Get(ctx context.Context, in *Xid, out *messages.Event) error {
	return h.StoreHandler.Get(ctx, in, out)
}

func (h *storeHandler) Del(ctx context.Context, in *Xid, out *GenericResponse) error {
	return h.StoreHandler.Del(ctx, in, out)
}

func (h *storeHandler) List(ctx context.Context, in *ListRequest, out *messages.Events) error {
	return h.StoreHandler.List(ctx, in, out)
}
