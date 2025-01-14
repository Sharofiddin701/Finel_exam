// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: super_admin_service.proto

package user_service

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
	SuperAdminService_Create_FullMethodName  = "/user_service.SuperAdminService/Create"
	SuperAdminService_GetByID_FullMethodName = "/user_service.SuperAdminService/GetByID"
	SuperAdminService_GetList_FullMethodName = "/user_service.SuperAdminService/GetList"
	SuperAdminService_Update_FullMethodName  = "/user_service.SuperAdminService/Update"
	SuperAdminService_Delete_FullMethodName  = "/user_service.SuperAdminService/Delete"
)

// SuperAdminServiceClient is the client API for SuperAdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SuperAdminServiceClient interface {
	Create(ctx context.Context, in *CreateSuperAdmin, opts ...grpc.CallOption) (*SuperAdmin, error)
	GetByID(ctx context.Context, in *SuperAdminPrimaryKey, opts ...grpc.CallOption) (*SuperAdmin, error)
	GetList(ctx context.Context, in *GetListSuperAdminRequest, opts ...grpc.CallOption) (*GetListSuperAdminResponse, error)
	Update(ctx context.Context, in *UpdateSuperAdmin, opts ...grpc.CallOption) (*SuperAdmin, error)
	Delete(ctx context.Context, in *SuperAdminPrimaryKey, opts ...grpc.CallOption) (*SuperAdminEmpty, error)
}

type superAdminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSuperAdminServiceClient(cc grpc.ClientConnInterface) SuperAdminServiceClient {
	return &superAdminServiceClient{cc}
}

func (c *superAdminServiceClient) Create(ctx context.Context, in *CreateSuperAdmin, opts ...grpc.CallOption) (*SuperAdmin, error) {
	out := new(SuperAdmin)
	err := c.cc.Invoke(ctx, SuperAdminService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *superAdminServiceClient) GetByID(ctx context.Context, in *SuperAdminPrimaryKey, opts ...grpc.CallOption) (*SuperAdmin, error) {
	out := new(SuperAdmin)
	err := c.cc.Invoke(ctx, SuperAdminService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *superAdminServiceClient) GetList(ctx context.Context, in *GetListSuperAdminRequest, opts ...grpc.CallOption) (*GetListSuperAdminResponse, error) {
	out := new(GetListSuperAdminResponse)
	err := c.cc.Invoke(ctx, SuperAdminService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *superAdminServiceClient) Update(ctx context.Context, in *UpdateSuperAdmin, opts ...grpc.CallOption) (*SuperAdmin, error) {
	out := new(SuperAdmin)
	err := c.cc.Invoke(ctx, SuperAdminService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *superAdminServiceClient) Delete(ctx context.Context, in *SuperAdminPrimaryKey, opts ...grpc.CallOption) (*SuperAdminEmpty, error) {
	out := new(SuperAdminEmpty)
	err := c.cc.Invoke(ctx, SuperAdminService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SuperAdminServiceServer is the server API for SuperAdminService service.
// All implementations must embed UnimplementedSuperAdminServiceServer
// for forward compatibility
type SuperAdminServiceServer interface {
	Create(context.Context, *CreateSuperAdmin) (*SuperAdmin, error)
	GetByID(context.Context, *SuperAdminPrimaryKey) (*SuperAdmin, error)
	GetList(context.Context, *GetListSuperAdminRequest) (*GetListSuperAdminResponse, error)
	Update(context.Context, *UpdateSuperAdmin) (*SuperAdmin, error)
	Delete(context.Context, *SuperAdminPrimaryKey) (*SuperAdminEmpty, error)
	mustEmbedUnimplementedSuperAdminServiceServer()
}

// UnimplementedSuperAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSuperAdminServiceServer struct {
}

func (UnimplementedSuperAdminServiceServer) Create(context.Context, *CreateSuperAdmin) (*SuperAdmin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSuperAdminServiceServer) GetByID(context.Context, *SuperAdminPrimaryKey) (*SuperAdmin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedSuperAdminServiceServer) GetList(context.Context, *GetListSuperAdminRequest) (*GetListSuperAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedSuperAdminServiceServer) Update(context.Context, *UpdateSuperAdmin) (*SuperAdmin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSuperAdminServiceServer) Delete(context.Context, *SuperAdminPrimaryKey) (*SuperAdminEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSuperAdminServiceServer) mustEmbedUnimplementedSuperAdminServiceServer() {}

// UnsafeSuperAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SuperAdminServiceServer will
// result in compilation errors.
type UnsafeSuperAdminServiceServer interface {
	mustEmbedUnimplementedSuperAdminServiceServer()
}

func RegisterSuperAdminServiceServer(s grpc.ServiceRegistrar, srv SuperAdminServiceServer) {
	s.RegisterService(&SuperAdminService_ServiceDesc, srv)
}

func _SuperAdminService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSuperAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuperAdminServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SuperAdminService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuperAdminServiceServer).Create(ctx, req.(*CreateSuperAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _SuperAdminService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuperAdminPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuperAdminServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SuperAdminService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuperAdminServiceServer).GetByID(ctx, req.(*SuperAdminPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _SuperAdminService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListSuperAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuperAdminServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SuperAdminService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuperAdminServiceServer).GetList(ctx, req.(*GetListSuperAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SuperAdminService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSuperAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuperAdminServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SuperAdminService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuperAdminServiceServer).Update(ctx, req.(*UpdateSuperAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _SuperAdminService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuperAdminPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuperAdminServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SuperAdminService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuperAdminServiceServer).Delete(ctx, req.(*SuperAdminPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// SuperAdminService_ServiceDesc is the grpc.ServiceDesc for SuperAdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SuperAdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_service.SuperAdminService",
	HandlerType: (*SuperAdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SuperAdminService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _SuperAdminService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _SuperAdminService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SuperAdminService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SuperAdminService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "super_admin_service.proto",
}
