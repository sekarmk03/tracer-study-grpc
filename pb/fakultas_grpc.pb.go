// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: fakultas.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	FakultasService_GetAllFakultas_FullMethodName    = "/tracer_study_grpc.FakultasService/GetAllFakultas"
	FakultasService_GetFakultasByKode_FullMethodName = "/tracer_study_grpc.FakultasService/GetFakultasByKode"
	FakultasService_CreateFakultas_FullMethodName    = "/tracer_study_grpc.FakultasService/CreateFakultas"
	FakultasService_UpdateFakultas_FullMethodName    = "/tracer_study_grpc.FakultasService/UpdateFakultas"
	FakultasService_DeleteFakultas_FullMethodName    = "/tracer_study_grpc.FakultasService/DeleteFakultas"
)

// FakultasServiceClient is the client API for FakultasService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FakultasServiceClient interface {
	GetAllFakultas(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllFakultasResponse, error)
	GetFakultasByKode(ctx context.Context, in *GetFakultasByKodeRequest, opts ...grpc.CallOption) (*GetFakultasResponse, error)
	CreateFakultas(ctx context.Context, in *Fakultas, opts ...grpc.CallOption) (*GetFakultasResponse, error)
	UpdateFakultas(ctx context.Context, in *Fakultas, opts ...grpc.CallOption) (*GetFakultasResponse, error)
	DeleteFakultas(ctx context.Context, in *GetFakultasByKodeRequest, opts ...grpc.CallOption) (*DeleteFakultasResponse, error)
}

type fakultasServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFakultasServiceClient(cc grpc.ClientConnInterface) FakultasServiceClient {
	return &fakultasServiceClient{cc}
}

func (c *fakultasServiceClient) GetAllFakultas(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllFakultasResponse, error) {
	out := new(GetAllFakultasResponse)
	err := c.cc.Invoke(ctx, FakultasService_GetAllFakultas_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fakultasServiceClient) GetFakultasByKode(ctx context.Context, in *GetFakultasByKodeRequest, opts ...grpc.CallOption) (*GetFakultasResponse, error) {
	out := new(GetFakultasResponse)
	err := c.cc.Invoke(ctx, FakultasService_GetFakultasByKode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fakultasServiceClient) CreateFakultas(ctx context.Context, in *Fakultas, opts ...grpc.CallOption) (*GetFakultasResponse, error) {
	out := new(GetFakultasResponse)
	err := c.cc.Invoke(ctx, FakultasService_CreateFakultas_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fakultasServiceClient) UpdateFakultas(ctx context.Context, in *Fakultas, opts ...grpc.CallOption) (*GetFakultasResponse, error) {
	out := new(GetFakultasResponse)
	err := c.cc.Invoke(ctx, FakultasService_UpdateFakultas_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fakultasServiceClient) DeleteFakultas(ctx context.Context, in *GetFakultasByKodeRequest, opts ...grpc.CallOption) (*DeleteFakultasResponse, error) {
	out := new(DeleteFakultasResponse)
	err := c.cc.Invoke(ctx, FakultasService_DeleteFakultas_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FakultasServiceServer is the server API for FakultasService service.
// All implementations must embed UnimplementedFakultasServiceServer
// for forward compatibility
type FakultasServiceServer interface {
	GetAllFakultas(context.Context, *emptypb.Empty) (*GetAllFakultasResponse, error)
	GetFakultasByKode(context.Context, *GetFakultasByKodeRequest) (*GetFakultasResponse, error)
	CreateFakultas(context.Context, *Fakultas) (*GetFakultasResponse, error)
	UpdateFakultas(context.Context, *Fakultas) (*GetFakultasResponse, error)
	DeleteFakultas(context.Context, *GetFakultasByKodeRequest) (*DeleteFakultasResponse, error)
	mustEmbedUnimplementedFakultasServiceServer()
}

// UnimplementedFakultasServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFakultasServiceServer struct {
}

func (UnimplementedFakultasServiceServer) GetAllFakultas(context.Context, *emptypb.Empty) (*GetAllFakultasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFakultas not implemented")
}
func (UnimplementedFakultasServiceServer) GetFakultasByKode(context.Context, *GetFakultasByKodeRequest) (*GetFakultasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFakultasByKode not implemented")
}
func (UnimplementedFakultasServiceServer) CreateFakultas(context.Context, *Fakultas) (*GetFakultasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFakultas not implemented")
}
func (UnimplementedFakultasServiceServer) UpdateFakultas(context.Context, *Fakultas) (*GetFakultasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFakultas not implemented")
}
func (UnimplementedFakultasServiceServer) DeleteFakultas(context.Context, *GetFakultasByKodeRequest) (*DeleteFakultasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFakultas not implemented")
}
func (UnimplementedFakultasServiceServer) mustEmbedUnimplementedFakultasServiceServer() {}

// UnsafeFakultasServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FakultasServiceServer will
// result in compilation errors.
type UnsafeFakultasServiceServer interface {
	mustEmbedUnimplementedFakultasServiceServer()
}

func RegisterFakultasServiceServer(s grpc.ServiceRegistrar, srv FakultasServiceServer) {
	s.RegisterService(&FakultasService_ServiceDesc, srv)
}

func _FakultasService_GetAllFakultas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FakultasServiceServer).GetAllFakultas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FakultasService_GetAllFakultas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FakultasServiceServer).GetAllFakultas(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FakultasService_GetFakultasByKode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFakultasByKodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FakultasServiceServer).GetFakultasByKode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FakultasService_GetFakultasByKode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FakultasServiceServer).GetFakultasByKode(ctx, req.(*GetFakultasByKodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FakultasService_CreateFakultas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Fakultas)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FakultasServiceServer).CreateFakultas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FakultasService_CreateFakultas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FakultasServiceServer).CreateFakultas(ctx, req.(*Fakultas))
	}
	return interceptor(ctx, in, info, handler)
}

func _FakultasService_UpdateFakultas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Fakultas)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FakultasServiceServer).UpdateFakultas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FakultasService_UpdateFakultas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FakultasServiceServer).UpdateFakultas(ctx, req.(*Fakultas))
	}
	return interceptor(ctx, in, info, handler)
}

func _FakultasService_DeleteFakultas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFakultasByKodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FakultasServiceServer).DeleteFakultas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FakultasService_DeleteFakultas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FakultasServiceServer).DeleteFakultas(ctx, req.(*GetFakultasByKodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FakultasService_ServiceDesc is the grpc.ServiceDesc for FakultasService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FakultasService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tracer_study_grpc.FakultasService",
	HandlerType: (*FakultasServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllFakultas",
			Handler:    _FakultasService_GetAllFakultas_Handler,
		},
		{
			MethodName: "GetFakultasByKode",
			Handler:    _FakultasService_GetFakultasByKode_Handler,
		},
		{
			MethodName: "CreateFakultas",
			Handler:    _FakultasService_CreateFakultas_Handler,
		},
		{
			MethodName: "UpdateFakultas",
			Handler:    _FakultasService_UpdateFakultas_Handler,
		},
		{
			MethodName: "DeleteFakultas",
			Handler:    _FakultasService_DeleteFakultas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fakultas.proto",
}
