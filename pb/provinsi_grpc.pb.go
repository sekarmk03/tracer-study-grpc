// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: provinsi.proto

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
	ProvinsiService_GetAllProvinsi_FullMethodName = "/tracer_study_grpc.ProvinsiService/GetAllProvinsi"
)

// ProvinsiServiceClient is the client API for ProvinsiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProvinsiServiceClient interface {
	GetAllProvinsi(ctx context.Context, in *EmptyProvinsiRequest, opts ...grpc.CallOption) (*GetAllProvinsiResponse, error)
}

type provinsiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProvinsiServiceClient(cc grpc.ClientConnInterface) ProvinsiServiceClient {
	return &provinsiServiceClient{cc}
}

func (c *provinsiServiceClient) GetAllProvinsi(ctx context.Context, in *EmptyProvinsiRequest, opts ...grpc.CallOption) (*GetAllProvinsiResponse, error) {
	out := new(GetAllProvinsiResponse)
	err := c.cc.Invoke(ctx, ProvinsiService_GetAllProvinsi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProvinsiServiceServer is the server API for ProvinsiService service.
// All implementations must embed UnimplementedProvinsiServiceServer
// for forward compatibility
type ProvinsiServiceServer interface {
	GetAllProvinsi(context.Context, *EmptyProvinsiRequest) (*GetAllProvinsiResponse, error)
	mustEmbedUnimplementedProvinsiServiceServer()
}

// UnimplementedProvinsiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProvinsiServiceServer struct {
}

func (UnimplementedProvinsiServiceServer) GetAllProvinsi(context.Context, *EmptyProvinsiRequest) (*GetAllProvinsiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProvinsi not implemented")
}
func (UnimplementedProvinsiServiceServer) mustEmbedUnimplementedProvinsiServiceServer() {}

// UnsafeProvinsiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProvinsiServiceServer will
// result in compilation errors.
type UnsafeProvinsiServiceServer interface {
	mustEmbedUnimplementedProvinsiServiceServer()
}

func RegisterProvinsiServiceServer(s grpc.ServiceRegistrar, srv ProvinsiServiceServer) {
	s.RegisterService(&ProvinsiService_ServiceDesc, srv)
}

func _ProvinsiService_GetAllProvinsi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyProvinsiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvinsiServiceServer).GetAllProvinsi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProvinsiService_GetAllProvinsi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvinsiServiceServer).GetAllProvinsi(ctx, req.(*EmptyProvinsiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProvinsiService_ServiceDesc is the grpc.ServiceDesc for ProvinsiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProvinsiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tracer_study_grpc.ProvinsiService",
	HandlerType: (*ProvinsiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllProvinsi",
			Handler:    _ProvinsiService_GetAllProvinsi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provinsi.proto",
}