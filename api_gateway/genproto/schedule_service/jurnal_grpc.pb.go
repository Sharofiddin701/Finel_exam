// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: jurnal.proto

package schedule_service

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
	JurnalService_Create_FullMethodName  = "/schedule_service.JurnalService/Create"
	JurnalService_GetByID_FullMethodName = "/schedule_service.JurnalService/GetByID"
	JurnalService_GetList_FullMethodName = "/schedule_service.JurnalService/GetList"
	JurnalService_Update_FullMethodName  = "/schedule_service.JurnalService/Update"
	JurnalService_Delete_FullMethodName  = "/schedule_service.JurnalService/Delete"
)

// JurnalServiceClient is the client API for JurnalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JurnalServiceClient interface {
	Create(ctx context.Context, in *CreateJurnal, opts ...grpc.CallOption) (*Jurnal, error)
	GetByID(ctx context.Context, in *JurnalPrimaryKey, opts ...grpc.CallOption) (*Jurnal, error)
	GetList(ctx context.Context, in *GetListJurnalRequest, opts ...grpc.CallOption) (*GetListJurnalResponse, error)
	Update(ctx context.Context, in *UpdateJurnal, opts ...grpc.CallOption) (*Jurnal, error)
	Delete(ctx context.Context, in *JurnalPrimaryKey, opts ...grpc.CallOption) (*JurnalEmpty, error)
}

type jurnalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJurnalServiceClient(cc grpc.ClientConnInterface) JurnalServiceClient {
	return &jurnalServiceClient{cc}
}

func (c *jurnalServiceClient) Create(ctx context.Context, in *CreateJurnal, opts ...grpc.CallOption) (*Jurnal, error) {
	out := new(Jurnal)
	err := c.cc.Invoke(ctx, JurnalService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jurnalServiceClient) GetByID(ctx context.Context, in *JurnalPrimaryKey, opts ...grpc.CallOption) (*Jurnal, error) {
	out := new(Jurnal)
	err := c.cc.Invoke(ctx, JurnalService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jurnalServiceClient) GetList(ctx context.Context, in *GetListJurnalRequest, opts ...grpc.CallOption) (*GetListJurnalResponse, error) {
	out := new(GetListJurnalResponse)
	err := c.cc.Invoke(ctx, JurnalService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jurnalServiceClient) Update(ctx context.Context, in *UpdateJurnal, opts ...grpc.CallOption) (*Jurnal, error) {
	out := new(Jurnal)
	err := c.cc.Invoke(ctx, JurnalService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jurnalServiceClient) Delete(ctx context.Context, in *JurnalPrimaryKey, opts ...grpc.CallOption) (*JurnalEmpty, error) {
	out := new(JurnalEmpty)
	err := c.cc.Invoke(ctx, JurnalService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JurnalServiceServer is the server API for JurnalService service.
// All implementations must embed UnimplementedJurnalServiceServer
// for forward compatibility
type JurnalServiceServer interface {
	Create(context.Context, *CreateJurnal) (*Jurnal, error)
	GetByID(context.Context, *JurnalPrimaryKey) (*Jurnal, error)
	GetList(context.Context, *GetListJurnalRequest) (*GetListJurnalResponse, error)
	Update(context.Context, *UpdateJurnal) (*Jurnal, error)
	Delete(context.Context, *JurnalPrimaryKey) (*JurnalEmpty, error)
	mustEmbedUnimplementedJurnalServiceServer()
}

// UnimplementedJurnalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJurnalServiceServer struct {
}

func (UnimplementedJurnalServiceServer) Create(context.Context, *CreateJurnal) (*Jurnal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedJurnalServiceServer) GetByID(context.Context, *JurnalPrimaryKey) (*Jurnal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedJurnalServiceServer) GetList(context.Context, *GetListJurnalRequest) (*GetListJurnalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedJurnalServiceServer) Update(context.Context, *UpdateJurnal) (*Jurnal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedJurnalServiceServer) Delete(context.Context, *JurnalPrimaryKey) (*JurnalEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedJurnalServiceServer) mustEmbedUnimplementedJurnalServiceServer() {}

// UnsafeJurnalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JurnalServiceServer will
// result in compilation errors.
type UnsafeJurnalServiceServer interface {
	mustEmbedUnimplementedJurnalServiceServer()
}

func RegisterJurnalServiceServer(s grpc.ServiceRegistrar, srv JurnalServiceServer) {
	s.RegisterService(&JurnalService_ServiceDesc, srv)
}

func _JurnalService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJurnal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JurnalServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JurnalService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JurnalServiceServer).Create(ctx, req.(*CreateJurnal))
	}
	return interceptor(ctx, in, info, handler)
}

func _JurnalService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JurnalPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JurnalServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JurnalService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JurnalServiceServer).GetByID(ctx, req.(*JurnalPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _JurnalService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListJurnalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JurnalServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JurnalService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JurnalServiceServer).GetList(ctx, req.(*GetListJurnalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JurnalService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateJurnal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JurnalServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JurnalService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JurnalServiceServer).Update(ctx, req.(*UpdateJurnal))
	}
	return interceptor(ctx, in, info, handler)
}

func _JurnalService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JurnalPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JurnalServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JurnalService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JurnalServiceServer).Delete(ctx, req.(*JurnalPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// JurnalService_ServiceDesc is the grpc.ServiceDesc for JurnalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JurnalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "schedule_service.JurnalService",
	HandlerType: (*JurnalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _JurnalService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _JurnalService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _JurnalService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _JurnalService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _JurnalService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jurnal.proto",
}