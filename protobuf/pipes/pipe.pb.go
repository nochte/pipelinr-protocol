// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pipe.proto

package pipes

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	messages "github.com/nochte/pipelinr-protocol/protobuf/messages"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GenericResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	OK                   bool     `protobuf:"varint,2,opt,name=OK,proto3" json:"OK,omitempty"`
	XId                  string   `protobuf:"bytes,3,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericResponse) Reset()         { *m = GenericResponse{} }
func (m *GenericResponse) String() string { return proto.CompactTextString(m) }
func (*GenericResponse) ProtoMessage()    {}
func (*GenericResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1979e1a5fdc07ed, []int{0}
}
func (m *GenericResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericResponse.Unmarshal(m, b)
}
func (m *GenericResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericResponse.Marshal(b, m, deterministic)
}
func (m *GenericResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericResponse.Merge(m, src)
}
func (m *GenericResponse) XXX_Size() int {
	return xxx_messageInfo_GenericResponse.Size(m)
}
func (m *GenericResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenericResponse proto.InternalMessageInfo

func (m *GenericResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GenericResponse) GetOK() bool {
	if m != nil {
		return m.OK
	}
	return false
}

func (m *GenericResponse) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

type GenericResponses struct {
	GenericResponses     []*GenericResponse `protobuf:"bytes,1,rep,name=GenericResponses,proto3" json:"GenericResponses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GenericResponses) Reset()         { *m = GenericResponses{} }
func (m *GenericResponses) String() string { return proto.CompactTextString(m) }
func (*GenericResponses) ProtoMessage()    {}
func (*GenericResponses) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1979e1a5fdc07ed, []int{1}
}
func (m *GenericResponses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericResponses.Unmarshal(m, b)
}
func (m *GenericResponses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericResponses.Marshal(b, m, deterministic)
}
func (m *GenericResponses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericResponses.Merge(m, src)
}
func (m *GenericResponses) XXX_Size() int {
	return xxx_messageInfo_GenericResponses.Size(m)
}
func (m *GenericResponses) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericResponses.DiscardUnknown(m)
}

var xxx_messageInfo_GenericResponses proto.InternalMessageInfo

func (m *GenericResponses) GetGenericResponses() []*GenericResponse {
	if m != nil {
		return m.GenericResponses
	}
	return nil
}

type Null struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Null) Reset()         { *m = Null{} }
func (m *Null) String() string { return proto.CompactTextString(m) }
func (*Null) ProtoMessage()    {}
func (*Null) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1979e1a5fdc07ed, []int{2}
}
func (m *Null) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Null.Unmarshal(m, b)
}
func (m *Null) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Null.Marshal(b, m, deterministic)
}
func (m *Null) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Null.Merge(m, src)
}
func (m *Null) XXX_Size() int {
	return xxx_messageInfo_Null.Size(m)
}
func (m *Null) XXX_DiscardUnknown() {
	xxx_messageInfo_Null.DiscardUnknown(m)
}

var xxx_messageInfo_Null proto.InternalMessageInfo

type ReceiveOptions struct {
	AutoAck              bool     `protobuf:"varint,1,opt,name=AutoAck,proto3" json:"AutoAck,omitempty"`
	Block                bool     `protobuf:"varint,2,opt,name=Block,proto3" json:"Block,omitempty"`
	Count                int32    `protobuf:"varint,3,opt,name=Count,proto3" json:"Count,omitempty"`
	Timeout              int64    `protobuf:"varint,4,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
	Pipe                 string   `protobuf:"bytes,5,opt,name=Pipe,proto3" json:"Pipe,omitempty"`
	RedeliveryTimeout    int64    `protobuf:"varint,6,opt,name=RedeliveryTimeout,proto3" json:"RedeliveryTimeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiveOptions) Reset()         { *m = ReceiveOptions{} }
func (m *ReceiveOptions) String() string { return proto.CompactTextString(m) }
func (*ReceiveOptions) ProtoMessage()    {}
func (*ReceiveOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1979e1a5fdc07ed, []int{3}
}
func (m *ReceiveOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveOptions.Unmarshal(m, b)
}
func (m *ReceiveOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveOptions.Marshal(b, m, deterministic)
}
func (m *ReceiveOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveOptions.Merge(m, src)
}
func (m *ReceiveOptions) XXX_Size() int {
	return xxx_messageInfo_ReceiveOptions.Size(m)
}
func (m *ReceiveOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveOptions proto.InternalMessageInfo

func (m *ReceiveOptions) GetAutoAck() bool {
	if m != nil {
		return m.AutoAck
	}
	return false
}

func (m *ReceiveOptions) GetBlock() bool {
	if m != nil {
		return m.Block
	}
	return false
}

func (m *ReceiveOptions) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReceiveOptions) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *ReceiveOptions) GetPipe() string {
	if m != nil {
		return m.Pipe
	}
	return ""
}

func (m *ReceiveOptions) GetRedeliveryTimeout() int64 {
	if m != nil {
		return m.RedeliveryTimeout
	}
	return 0
}

type Xid struct {
	XId                  string   `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Xid) Reset()         { *m = Xid{} }
func (m *Xid) String() string { return proto.CompactTextString(m) }
func (*Xid) ProtoMessage()    {}
func (*Xid) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1979e1a5fdc07ed, []int{4}
}
func (m *Xid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Xid.Unmarshal(m, b)
}
func (m *Xid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Xid.Marshal(b, m, deterministic)
}
func (m *Xid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Xid.Merge(m, src)
}
func (m *Xid) XXX_Size() int {
	return xxx_messageInfo_Xid.Size(m)
}
func (m *Xid) XXX_DiscardUnknown() {
	xxx_messageInfo_Xid.DiscardUnknown(m)
}

var xxx_messageInfo_Xid proto.InternalMessageInfo

func (m *Xid) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

type CompleteRequest struct {
	XId                  string   `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	Step                 string   `protobuf:"bytes,2,opt,name=step,proto3" json:"step,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompleteRequest) Reset()         { *m = CompleteRequest{} }
func (m *CompleteRequest) String() string { return proto.CompactTextString(m) }
func (*CompleteRequest) ProtoMessage()    {}
func (*CompleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1979e1a5fdc07ed, []int{5}
}
func (m *CompleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteRequest.Unmarshal(m, b)
}
func (m *CompleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteRequest.Marshal(b, m, deterministic)
}
func (m *CompleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteRequest.Merge(m, src)
}
func (m *CompleteRequest) XXX_Size() int {
	return xxx_messageInfo_CompleteRequest.Size(m)
}
func (m *CompleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteRequest proto.InternalMessageInfo

func (m *CompleteRequest) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *CompleteRequest) GetStep() string {
	if m != nil {
		return m.Step
	}
	return ""
}

type RouteLogRequest struct {
	XId                  string             `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	Log                  *messages.RouteLog `protobuf:"bytes,2,opt,name=Log,proto3" json:"Log,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RouteLogRequest) Reset()         { *m = RouteLogRequest{} }
func (m *RouteLogRequest) String() string { return proto.CompactTextString(m) }
func (*RouteLogRequest) ProtoMessage()    {}
func (*RouteLogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1979e1a5fdc07ed, []int{6}
}
func (m *RouteLogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteLogRequest.Unmarshal(m, b)
}
func (m *RouteLogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteLogRequest.Marshal(b, m, deterministic)
}
func (m *RouteLogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteLogRequest.Merge(m, src)
}
func (m *RouteLogRequest) XXX_Size() int {
	return xxx_messageInfo_RouteLogRequest.Size(m)
}
func (m *RouteLogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteLogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RouteLogRequest proto.InternalMessageInfo

func (m *RouteLogRequest) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *RouteLogRequest) GetLog() *messages.RouteLog {
	if m != nil {
		return m.Log
	}
	return nil
}

type AddStepsRequest struct {
	XId                  string   `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	After                string   `protobuf:"bytes,2,opt,name=After,proto3" json:"After,omitempty"`
	NewSteps             []string `protobuf:"bytes,3,rep,name=NewSteps,proto3" json:"NewSteps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddStepsRequest) Reset()         { *m = AddStepsRequest{} }
func (m *AddStepsRequest) String() string { return proto.CompactTextString(m) }
func (*AddStepsRequest) ProtoMessage()    {}
func (*AddStepsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1979e1a5fdc07ed, []int{7}
}
func (m *AddStepsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddStepsRequest.Unmarshal(m, b)
}
func (m *AddStepsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddStepsRequest.Marshal(b, m, deterministic)
}
func (m *AddStepsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddStepsRequest.Merge(m, src)
}
func (m *AddStepsRequest) XXX_Size() int {
	return xxx_messageInfo_AddStepsRequest.Size(m)
}
func (m *AddStepsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddStepsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddStepsRequest proto.InternalMessageInfo

func (m *AddStepsRequest) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *AddStepsRequest) GetAfter() string {
	if m != nil {
		return m.After
	}
	return ""
}

func (m *AddStepsRequest) GetNewSteps() []string {
	if m != nil {
		return m.NewSteps
	}
	return nil
}

type Decoration struct {
	XId string `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	Key string `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
	// google.protobuf.Value Value = 3;
	Value                string   `protobuf:"bytes,4,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Decoration) Reset()         { *m = Decoration{} }
func (m *Decoration) String() string { return proto.CompactTextString(m) }
func (*Decoration) ProtoMessage()    {}
func (*Decoration) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1979e1a5fdc07ed, []int{8}
}
func (m *Decoration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Decoration.Unmarshal(m, b)
}
func (m *Decoration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Decoration.Marshal(b, m, deterministic)
}
func (m *Decoration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Decoration.Merge(m, src)
}
func (m *Decoration) XXX_Size() int {
	return xxx_messageInfo_Decoration.Size(m)
}
func (m *Decoration) XXX_DiscardUnknown() {
	xxx_messageInfo_Decoration.DiscardUnknown(m)
}

var xxx_messageInfo_Decoration proto.InternalMessageInfo

func (m *Decoration) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *Decoration) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Decoration) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Decorations struct {
	XId                  string        `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	Decorations          []*Decoration `protobuf:"bytes,2,rep,name=Decorations,proto3" json:"Decorations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Decorations) Reset()         { *m = Decorations{} }
func (m *Decorations) String() string { return proto.CompactTextString(m) }
func (*Decorations) ProtoMessage()    {}
func (*Decorations) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1979e1a5fdc07ed, []int{9}
}
func (m *Decorations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Decorations.Unmarshal(m, b)
}
func (m *Decorations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Decorations.Marshal(b, m, deterministic)
}
func (m *Decorations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Decorations.Merge(m, src)
}
func (m *Decorations) XXX_Size() int {
	return xxx_messageInfo_Decorations.Size(m)
}
func (m *Decorations) XXX_DiscardUnknown() {
	xxx_messageInfo_Decorations.DiscardUnknown(m)
}

var xxx_messageInfo_Decorations proto.InternalMessageInfo

func (m *Decorations) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *Decorations) GetDecorations() []*Decoration {
	if m != nil {
		return m.Decorations
	}
	return nil
}

type GetDecorationRequest struct {
	XId                  string   `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	Keys                 []string `protobuf:"bytes,2,rep,name=Keys,proto3" json:"Keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDecorationRequest) Reset()         { *m = GetDecorationRequest{} }
func (m *GetDecorationRequest) String() string { return proto.CompactTextString(m) }
func (*GetDecorationRequest) ProtoMessage()    {}
func (*GetDecorationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1979e1a5fdc07ed, []int{10}
}
func (m *GetDecorationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDecorationRequest.Unmarshal(m, b)
}
func (m *GetDecorationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDecorationRequest.Marshal(b, m, deterministic)
}
func (m *GetDecorationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDecorationRequest.Merge(m, src)
}
func (m *GetDecorationRequest) XXX_Size() int {
	return xxx_messageInfo_GetDecorationRequest.Size(m)
}
func (m *GetDecorationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDecorationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDecorationRequest proto.InternalMessageInfo

func (m *GetDecorationRequest) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *GetDecorationRequest) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type ListRequest struct {
	UserId               string                 `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	StartTime            *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime              *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	Limit                int64                  `protobuf:"varint,4,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset               int64                  `protobuf:"varint,5,opt,name=Offset,proto3" json:"Offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1979e1a5fdc07ed, []int{11}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ListRequest) GetStartTime() *timestamppb.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ListRequest) GetEndTime() *timestamppb.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *ListRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func init() {
	proto.RegisterType((*GenericResponse)(nil), "pipes.GenericResponse")
	proto.RegisterType((*GenericResponses)(nil), "pipes.GenericResponses")
	proto.RegisterType((*Null)(nil), "pipes.Null")
	proto.RegisterType((*ReceiveOptions)(nil), "pipes.ReceiveOptions")
	proto.RegisterType((*Xid)(nil), "pipes.Xid")
	proto.RegisterType((*CompleteRequest)(nil), "pipes.CompleteRequest")
	proto.RegisterType((*RouteLogRequest)(nil), "pipes.RouteLogRequest")
	proto.RegisterType((*AddStepsRequest)(nil), "pipes.AddStepsRequest")
	proto.RegisterType((*Decoration)(nil), "pipes.Decoration")
	proto.RegisterType((*Decorations)(nil), "pipes.Decorations")
	proto.RegisterType((*GetDecorationRequest)(nil), "pipes.GetDecorationRequest")
	proto.RegisterType((*ListRequest)(nil), "pipes.ListRequest")
}

func init() { proto.RegisterFile("pipe.proto", fileDescriptor_d1979e1a5fdc07ed) }

var fileDescriptor_d1979e1a5fdc07ed = []byte{
	// 784 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x8e, 0xe3, 0xa4, 0x9b, 0x9c, 0x48, 0x4d, 0x77, 0x54, 0x8a, 0x65, 0x2e, 0x88, 0x46, 0x5c,
	0x44, 0x02, 0xdc, 0x25, 0xbb, 0x82, 0x45, 0xc0, 0x45, 0xda, 0xad, 0xca, 0x2a, 0x61, 0x83, 0x26,
	0x65, 0xb5, 0xe2, 0x06, 0xa5, 0xf6, 0x49, 0x76, 0x54, 0xdb, 0x63, 0x3c, 0xe3, 0xa0, 0xbe, 0x02,
	0x8f, 0xc3, 0x2b, 0xf0, 0x04, 0xbc, 0x11, 0x9a, 0xb1, 0x9d, 0x3f, 0x27, 0x41, 0xec, 0xdd, 0x7c,
	0x3e, 0xe7, 0xfb, 0xe6, 0xcc, 0x99, 0xf3, 0x8d, 0x01, 0x12, 0x9e, 0xa0, 0x97, 0xa4, 0x42, 0x09,
	0xd2, 0xd4, 0x6b, 0xe9, 0x9e, 0x47, 0x28, 0xe5, 0x6c, 0x81, 0x18, 0x2f, 0x31, 0x14, 0x49, 0x1e,
	0x74, 0x3b, 0xb8, 0xc4, 0x58, 0x15, 0xe0, 0xd3, 0x85, 0x10, 0x8b, 0x10, 0x2f, 0x0d, 0xba, 0xcf,
	0xe6, 0x97, 0x8a, 0x47, 0x28, 0xd5, 0x2c, 0x2a, 0xb2, 0xe9, 0x18, 0xba, 0xb7, 0x18, 0x63, 0xca,
	0x7d, 0x86, 0x32, 0x11, 0xb1, 0x44, 0xe2, 0xc0, 0x93, 0x9f, 0x72, 0x61, 0xc7, 0xea, 0x59, 0xfd,
	0x36, 0x2b, 0x21, 0x39, 0x85, 0xfa, 0x64, 0xe4, 0xd4, 0x7b, 0x56, 0xbf, 0xc5, 0xea, 0x93, 0x11,
	0xe9, 0x82, 0xfd, 0x1b, 0x0f, 0x1c, 0xdb, 0x64, 0xd5, 0x5f, 0x07, 0xf4, 0x2d, 0x9c, 0xed, 0xa8,
	0x49, 0x72, 0x55, 0xfd, 0xe6, 0x58, 0x3d, 0xbb, 0xdf, 0x19, 0x5c, 0x78, 0xe6, 0x1c, 0xde, 0x4e,
	0x98, 0x55, 0xf2, 0xe9, 0x09, 0x34, 0xde, 0x64, 0x61, 0x48, 0xff, 0xb2, 0xe0, 0x94, 0xa1, 0x8f,
	0x7c, 0x89, 0x93, 0x44, 0x71, 0x11, 0x4b, 0x5d, 0xed, 0x30, 0x53, 0x62, 0xe8, 0x3f, 0x98, 0x6a,
	0x5b, 0xac, 0x84, 0xe4, 0x1c, 0x9a, 0x57, 0xa1, 0xf0, 0x1f, 0x8a, 0x82, 0x73, 0xa0, 0xbf, 0x5e,
	0x8b, 0x2c, 0x56, 0xa6, 0xea, 0x26, 0xcb, 0x81, 0x56, 0xb9, 0xe3, 0x11, 0x8a, 0x4c, 0x39, 0x8d,
	0x9e, 0xd5, 0xb7, 0x59, 0x09, 0x09, 0x81, 0xc6, 0xcf, 0x3c, 0x41, 0xa7, 0x69, 0x0e, 0x69, 0xd6,
	0xe4, 0x0b, 0x78, 0xca, 0x30, 0xc0, 0x90, 0x2f, 0x31, 0x7d, 0x2c, 0x79, 0x27, 0x86, 0x57, 0x0d,
	0xd0, 0x0b, 0xb0, 0xdf, 0xf1, 0xa0, 0x6c, 0x96, 0xb5, 0x6a, 0xd6, 0xd7, 0xd0, 0xbd, 0x16, 0x51,
	0x12, 0xa2, 0x42, 0x86, 0xbf, 0x67, 0x28, 0x55, 0x25, 0x47, 0xef, 0x2e, 0x15, 0x26, 0xe6, 0x08,
	0x6d, 0x66, 0xd6, 0xf4, 0x47, 0xe8, 0x32, 0x91, 0x29, 0x1c, 0x8b, 0xc5, 0x41, 0xde, 0x67, 0x60,
	0x8f, 0xc5, 0xc2, 0xd0, 0x3a, 0x03, 0xe2, 0x15, 0x83, 0x22, 0xbd, 0x15, 0x51, 0x87, 0xe9, 0x1d,
	0x74, 0x87, 0x41, 0x30, 0x55, 0x98, 0xc8, 0x83, 0x4a, 0xe7, 0xd0, 0x1c, 0xce, 0x15, 0xa6, 0x45,
	0x09, 0x39, 0x20, 0x2e, 0xb4, 0xde, 0xe0, 0x1f, 0x86, 0xe9, 0xd8, 0x3d, 0xbb, 0xdf, 0x66, 0x2b,
	0x4c, 0x6f, 0x00, 0x5e, 0xa1, 0x2f, 0xd2, 0x99, 0xbe, 0xa0, 0xaa, 0xe0, 0x19, 0xd8, 0x23, 0x7c,
	0x2c, 0xe4, 0xf4, 0x52, 0x6f, 0xf1, 0x76, 0x16, 0x66, 0x68, 0x5a, 0xdf, 0x66, 0x39, 0xa0, 0x53,
	0xe8, 0xac, 0x65, 0x64, 0x55, 0xe7, 0xf9, 0x56, 0xdc, 0xa9, 0x9b, 0x91, 0x7a, 0x5a, 0x8c, 0xd4,
	0x3a, 0xc2, 0x36, 0xb3, 0xe8, 0x77, 0x70, 0x7e, 0x8b, 0x6a, 0x23, 0x7a, 0xa4, 0xf1, 0x23, 0x7c,
	0xcc, 0x65, 0xdb, 0xcc, 0xac, 0xe9, 0xdf, 0x16, 0x74, 0xc6, 0x5c, 0xaa, 0x92, 0x74, 0x01, 0x27,
	0xbf, 0x48, 0x4c, 0x5f, 0x97, 0xbc, 0x02, 0x91, 0x97, 0xd0, 0x9e, 0xaa, 0x59, 0xaa, 0xf4, 0x00,
	0x14, 0x57, 0xe0, 0x7a, 0xb9, 0x11, 0xbd, 0xd2, 0x88, 0xde, 0x5d, 0x69, 0x44, 0xb6, 0x4e, 0x26,
	0x2f, 0xe0, 0xc9, 0x4d, 0x1c, 0x18, 0x9e, 0xfd, 0x9f, 0xbc, 0x32, 0x55, 0xf7, 0x6f, 0xcc, 0x23,
	0x5e, 0x8e, 0x6e, 0x0e, 0x74, 0x75, 0x93, 0xf9, 0x5c, 0xa2, 0x32, 0xa3, 0x6b, 0xb3, 0x02, 0x0d,
	0xfe, 0x69, 0xe4, 0x13, 0x4d, 0x5e, 0x02, 0x4c, 0x55, 0x8a, 0xb3, 0x68, 0x8a, 0x71, 0x40, 0x9c,
	0xf5, 0x90, 0x14, 0x76, 0xbf, 0xc9, 0x9f, 0x15, 0x17, 0x8a, 0x9e, 0xbe, 0xe3, 0x01, 0xad, 0xf5,
	0xad, 0x67, 0x16, 0xf1, 0xa0, 0xf1, 0x7f, 0x38, 0xe4, 0x2b, 0x68, 0x30, 0xf4, 0x97, 0xe4, 0xa3,
	0xe2, 0xeb, 0xb6, 0x85, 0xdd, 0xb3, 0xb5, 0xcc, 0x8d, 0x7e, 0xbb, 0x24, 0xad, 0xad, 0x8b, 0x3b,
	0x46, 0xec, 0xee, 0x10, 0x69, 0xed, 0x99, 0x45, 0xbe, 0x01, 0x5b, 0xbb, 0xbf, 0x7c, 0x5c, 0x76,
	0x2c, 0xe6, 0x1e, 0x78, 0x74, 0x68, 0x8d, 0x7c, 0x0f, 0xad, 0x32, 0xf9, 0x03, 0xd8, 0x3f, 0x40,
	0x7b, 0x98, 0x24, 0x18, 0x07, 0x63, 0xb1, 0x58, 0xd1, 0x77, 0x7c, 0x7a, 0x7c, 0xf3, 0xd2, 0x8a,
	0x2b, 0xf6, 0x8e, 0x37, 0x8f, 0xb0, 0xbf, 0x85, 0x56, 0x31, 0xd3, 0x48, 0x48, 0xc5, 0x02, 0xd2,
	0xfd, 0x78, 0x3f, 0x53, 0x37, 0xfa, 0x1a, 0x4e, 0xb7, 0x1c, 0x21, 0xc9, 0x27, 0xab, 0xe4, 0xaa,
	0x51, 0xdc, 0x3d, 0xea, 0xb4, 0x36, 0xf8, 0xd3, 0x82, 0xe6, 0x54, 0x89, 0x14, 0xf5, 0xc3, 0x73,
	0x8b, 0x8a, 0x6c, 0xdc, 0xff, 0x9e, 0x5b, 0x22, 0x9f, 0x83, 0xfd, 0x0a, 0xc3, 0xad, 0xac, 0xc3,
	0x87, 0xf3, 0xa0, 0xa1, 0x5d, 0xb7, 0x3a, 0xd8, 0x86, 0x05, 0xf7, 0x8d, 0xce, 0xd5, 0x8b, 0x5f,
	0x07, 0x0b, 0xae, 0xde, 0x67, 0xf7, 0x9e, 0x2f, 0xa2, 0xcb, 0x58, 0xf8, 0xef, 0x15, 0x5e, 0x6a,
	0x6a, 0xc8, 0xe3, 0xf4, 0x4b, 0x63, 0x24, 0x5f, 0x84, 0xeb, 0x5f, 0xa2, 0x51, 0xbd, 0x3f, 0x31,
	0xf8, 0xf9, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xeb, 0x1e, 0xb4, 0x68, 0x07, 0x00, 0x00,
}
