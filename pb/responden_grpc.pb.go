// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: responden.proto

package pb

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
	RespondenService_GetAllResponden_FullMethodName = "/tracer_study_grpc.RespondenService/GetAllResponden"
)

// RespondenServiceClient is the client API for RespondenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RespondenServiceClient interface {
	GetAllResponden(ctx context.Context, in *EmptyRespondenRequest, opts ...grpc.CallOption) (*GetAllRespondenResponse, error)
}

type respondenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRespondenServiceClient(cc grpc.ClientConnInterface) RespondenServiceClient {
	return &respondenServiceClient{cc}
}

func (c *respondenServiceClient) GetAllResponden(ctx context.Context, in *EmptyRespondenRequest, opts ...grpc.CallOption) (*GetAllRespondenResponse, error) {
	out := new(GetAllRespondenResponse)
	err := c.cc.Invoke(ctx, RespondenService_GetAllResponden_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RespondenServiceServer is the server API for RespondenService service.
// All implementations must embed UnimplementedRespondenServiceServer
// for forward compatibility
type RespondenServiceServer interface {
	GetAllResponden(context.Context, *EmptyRespondenRequest) (*GetAllRespondenResponse, error)
	mustEmbedUnimplementedRespondenServiceServer()
}

// UnimplementedRespondenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRespondenServiceServer struct {
}

func (UnimplementedRespondenServiceServer) GetAllResponden(context.Context, *EmptyRespondenRequest) (*GetAllRespondenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllResponden not implemented")
}
func (UnimplementedRespondenServiceServer) mustEmbedUnimplementedRespondenServiceServer() {}

// UnsafeRespondenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RespondenServiceServer will
// result in compilation errors.
type UnsafeRespondenServiceServer interface {
	mustEmbedUnimplementedRespondenServiceServer()
}

func RegisterRespondenServiceServer(s grpc.ServiceRegistrar, srv RespondenServiceServer) {
	s.RegisterService(&RespondenService_ServiceDesc, srv)
}

func _RespondenService_GetAllResponden_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRespondenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RespondenServiceServer).GetAllResponden(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RespondenService_GetAllResponden_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RespondenServiceServer).GetAllResponden(ctx, req.(*EmptyRespondenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RespondenService_ServiceDesc is the grpc.ServiceDesc for RespondenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RespondenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tracer_study_grpc.RespondenService",
	HandlerType: (*RespondenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllResponden",
			Handler:    _RespondenService_GetAllResponden_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "responden.proto",
}
