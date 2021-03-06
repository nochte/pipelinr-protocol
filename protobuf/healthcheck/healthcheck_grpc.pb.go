// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package healthcheck

import (
	context "context"
	pipes "github.com/nochte/pipelinr-protocol/protobuf/pipes"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HealthcheckClient is the client API for Healthcheck service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthcheckClient interface {
	IsReady(ctx context.Context, in *pipes.Null, opts ...grpc.CallOption) (*pipes.GenericResponse, error)
	IsHealthy(ctx context.Context, in *pipes.Null, opts ...grpc.CallOption) (*pipes.GenericResponse, error)
}

type healthcheckClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthcheckClient(cc grpc.ClientConnInterface) HealthcheckClient {
	return &healthcheckClient{cc}
}

func (c *healthcheckClient) IsReady(ctx context.Context, in *pipes.Null, opts ...grpc.CallOption) (*pipes.GenericResponse, error) {
	out := new(pipes.GenericResponse)
	err := c.cc.Invoke(ctx, "/healthcheck.Healthcheck/IsReady", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthcheckClient) IsHealthy(ctx context.Context, in *pipes.Null, opts ...grpc.CallOption) (*pipes.GenericResponse, error) {
	out := new(pipes.GenericResponse)
	err := c.cc.Invoke(ctx, "/healthcheck.Healthcheck/IsHealthy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthcheckServer is the server API for Healthcheck service.
// All implementations should embed UnimplementedHealthcheckServer
// for forward compatibility
type HealthcheckServer interface {
	IsReady(context.Context, *pipes.Null) (*pipes.GenericResponse, error)
	IsHealthy(context.Context, *pipes.Null) (*pipes.GenericResponse, error)
}

// UnimplementedHealthcheckServer should be embedded to have forward compatible implementations.
type UnimplementedHealthcheckServer struct {
}

func (UnimplementedHealthcheckServer) IsReady(context.Context, *pipes.Null) (*pipes.GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReady not implemented")
}
func (UnimplementedHealthcheckServer) IsHealthy(context.Context, *pipes.Null) (*pipes.GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsHealthy not implemented")
}

// UnsafeHealthcheckServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthcheckServer will
// result in compilation errors.
type UnsafeHealthcheckServer interface {
	mustEmbedUnimplementedHealthcheckServer()
}

func RegisterHealthcheckServer(s grpc.ServiceRegistrar, srv HealthcheckServer) {
	s.RegisterService(&Healthcheck_ServiceDesc, srv)
}

func _Healthcheck_IsReady_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pipes.Null)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthcheckServer).IsReady(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/healthcheck.Healthcheck/IsReady",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthcheckServer).IsReady(ctx, req.(*pipes.Null))
	}
	return interceptor(ctx, in, info, handler)
}

func _Healthcheck_IsHealthy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pipes.Null)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthcheckServer).IsHealthy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/healthcheck.Healthcheck/IsHealthy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthcheckServer).IsHealthy(ctx, req.(*pipes.Null))
	}
	return interceptor(ctx, in, info, handler)
}

// Healthcheck_ServiceDesc is the grpc.ServiceDesc for Healthcheck service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Healthcheck_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "healthcheck.Healthcheck",
	HandlerType: (*HealthcheckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReady",
			Handler:    _Healthcheck_IsReady_Handler,
		},
		{
			MethodName: "IsHealthy",
			Handler:    _Healthcheck_IsHealthy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "healthcheck.proto",
}
