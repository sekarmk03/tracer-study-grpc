// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: kabkota.proto

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
	KabKotaService_GetAllKabKota_FullMethodName = "/tracer_study_grpc.KabKotaService/GetAllKabKota"
)

// KabKotaServiceClient is the client API for KabKotaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KabKotaServiceClient interface {
	GetAllKabKota(ctx context.Context, in *EmptyKabKotaRequest, opts ...grpc.CallOption) (*GetAllKabKotaResponse, error)
}

type kabKotaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKabKotaServiceClient(cc grpc.ClientConnInterface) KabKotaServiceClient {
	return &kabKotaServiceClient{cc}
}

func (c *kabKotaServiceClient) GetAllKabKota(ctx context.Context, in *EmptyKabKotaRequest, opts ...grpc.CallOption) (*GetAllKabKotaResponse, error) {
	out := new(GetAllKabKotaResponse)
	err := c.cc.Invoke(ctx, KabKotaService_GetAllKabKota_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KabKotaServiceServer is the server API for KabKotaService service.
// All implementations must embed UnimplementedKabKotaServiceServer
// for forward compatibility
type KabKotaServiceServer interface {
	GetAllKabKota(context.Context, *EmptyKabKotaRequest) (*GetAllKabKotaResponse, error)
	mustEmbedUnimplementedKabKotaServiceServer()
}

// UnimplementedKabKotaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKabKotaServiceServer struct {
}

func (UnimplementedKabKotaServiceServer) GetAllKabKota(context.Context, *EmptyKabKotaRequest) (*GetAllKabKotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllKabKota not implemented")
}
func (UnimplementedKabKotaServiceServer) mustEmbedUnimplementedKabKotaServiceServer() {}

// UnsafeKabKotaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KabKotaServiceServer will
// result in compilation errors.
type UnsafeKabKotaServiceServer interface {
	mustEmbedUnimplementedKabKotaServiceServer()
}

func RegisterKabKotaServiceServer(s grpc.ServiceRegistrar, srv KabKotaServiceServer) {
	s.RegisterService(&KabKotaService_ServiceDesc, srv)
}

func _KabKotaService_GetAllKabKota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyKabKotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KabKotaServiceServer).GetAllKabKota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KabKotaService_GetAllKabKota_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KabKotaServiceServer).GetAllKabKota(ctx, req.(*EmptyKabKotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KabKotaService_ServiceDesc is the grpc.ServiceDesc for KabKotaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KabKotaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tracer_study_grpc.KabKotaService",
	HandlerType: (*KabKotaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllKabKota",
			Handler:    _KabKotaService_GetAllKabKota_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kabkota.proto",
}