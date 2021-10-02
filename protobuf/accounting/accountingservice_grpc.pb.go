// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package accounting

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

// AccountingClient is the client API for Accounting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountingClient interface {
	GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*GetReportResponse, error)
}

type accountingClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountingClient(cc grpc.ClientConnInterface) AccountingClient {
	return &accountingClient{cc}
}

func (c *accountingClient) GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*GetReportResponse, error) {
	out := new(GetReportResponse)
	err := c.cc.Invoke(ctx, "/accounting.Accounting/GetReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountingServer is the server API for Accounting service.
// All implementations should embed UnimplementedAccountingServer
// for forward compatibility
type AccountingServer interface {
	GetReport(context.Context, *GetReportRequest) (*GetReportResponse, error)
}

// UnimplementedAccountingServer should be embedded to have forward compatible implementations.
type UnimplementedAccountingServer struct {
}

func (UnimplementedAccountingServer) GetReport(context.Context, *GetReportRequest) (*GetReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReport not implemented")
}

// UnsafeAccountingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountingServer will
// result in compilation errors.
type UnsafeAccountingServer interface {
	mustEmbedUnimplementedAccountingServer()
}

func RegisterAccountingServer(s grpc.ServiceRegistrar, srv AccountingServer) {
	s.RegisterService(&Accounting_ServiceDesc, srv)
}

func _Accounting_GetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountingServer).GetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounting.Accounting/GetReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountingServer).GetReport(ctx, req.(*GetReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Accounting_ServiceDesc is the grpc.ServiceDesc for Accounting service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Accounting_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accounting.Accounting",
	HandlerType: (*AccountingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReport",
			Handler:    _Accounting_GetReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accountingservice.proto",
}
