// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/user/v1/services/user_reg_email.service.v1.proto

package userservicev1

import (
	context "context"
	resources "github.com/ikaiguang/go-srv-user/api/user/v1/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SrvUserRegEmailClient is the client API for SrvUserRegEmail service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SrvUserRegEmailClient interface {
	// GetById get by id
	GetById(ctx context.Context, in *resources.UserRegEmailIdReq, opts ...grpc.CallOption) (*resources.UserRegEmail, error)
	// List list
	List(ctx context.Context, in *resources.UserRegEmailListReq, opts ...grpc.CallOption) (*resources.UserRegEmailListResp, error)
	// Create create
	Create(ctx context.Context, in *resources.UserRegEmailSaveReq, opts ...grpc.CallOption) (*resources.UserRegEmail, error)
	// Update update
	Update(ctx context.Context, in *resources.UserRegEmailSaveReq, opts ...grpc.CallOption) (*resources.UserRegEmail, error)
	// Delete delete
	Delete(ctx context.Context, in *resources.UserRegEmailIdReq, opts ...grpc.CallOption) (*resources.UserRegEmailProcessResult, error)
	// Destroy destroy
	Destroy(ctx context.Context, in *resources.UserRegEmailIdsReq, opts ...grpc.CallOption) (*resources.UserRegEmailProcessResult, error)
}

type srvUserRegEmailClient struct {
	cc grpc.ClientConnInterface
}

func NewSrvUserRegEmailClient(cc grpc.ClientConnInterface) SrvUserRegEmailClient {
	return &srvUserRegEmailClient{cc}
}

func (c *srvUserRegEmailClient) GetById(ctx context.Context, in *resources.UserRegEmailIdReq, opts ...grpc.CallOption) (*resources.UserRegEmail, error) {
	out := new(resources.UserRegEmail)
	err := c.cc.Invoke(ctx, "/user.api.user.userservicev1.SrvUserRegEmail/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUserRegEmailClient) List(ctx context.Context, in *resources.UserRegEmailListReq, opts ...grpc.CallOption) (*resources.UserRegEmailListResp, error) {
	out := new(resources.UserRegEmailListResp)
	err := c.cc.Invoke(ctx, "/user.api.user.userservicev1.SrvUserRegEmail/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUserRegEmailClient) Create(ctx context.Context, in *resources.UserRegEmailSaveReq, opts ...grpc.CallOption) (*resources.UserRegEmail, error) {
	out := new(resources.UserRegEmail)
	err := c.cc.Invoke(ctx, "/user.api.user.userservicev1.SrvUserRegEmail/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUserRegEmailClient) Update(ctx context.Context, in *resources.UserRegEmailSaveReq, opts ...grpc.CallOption) (*resources.UserRegEmail, error) {
	out := new(resources.UserRegEmail)
	err := c.cc.Invoke(ctx, "/user.api.user.userservicev1.SrvUserRegEmail/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUserRegEmailClient) Delete(ctx context.Context, in *resources.UserRegEmailIdReq, opts ...grpc.CallOption) (*resources.UserRegEmailProcessResult, error) {
	out := new(resources.UserRegEmailProcessResult)
	err := c.cc.Invoke(ctx, "/user.api.user.userservicev1.SrvUserRegEmail/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUserRegEmailClient) Destroy(ctx context.Context, in *resources.UserRegEmailIdsReq, opts ...grpc.CallOption) (*resources.UserRegEmailProcessResult, error) {
	out := new(resources.UserRegEmailProcessResult)
	err := c.cc.Invoke(ctx, "/user.api.user.userservicev1.SrvUserRegEmail/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SrvUserRegEmailServer is the server API for SrvUserRegEmail service.
// All implementations must embed UnimplementedSrvUserRegEmailServer
// for forward compatibility
type SrvUserRegEmailServer interface {
	// GetById get by id
	GetById(context.Context, *resources.UserRegEmailIdReq) (*resources.UserRegEmail, error)
	// List list
	List(context.Context, *resources.UserRegEmailListReq) (*resources.UserRegEmailListResp, error)
	// Create create
	Create(context.Context, *resources.UserRegEmailSaveReq) (*resources.UserRegEmail, error)
	// Update update
	Update(context.Context, *resources.UserRegEmailSaveReq) (*resources.UserRegEmail, error)
	// Delete delete
	Delete(context.Context, *resources.UserRegEmailIdReq) (*resources.UserRegEmailProcessResult, error)
	// Destroy destroy
	Destroy(context.Context, *resources.UserRegEmailIdsReq) (*resources.UserRegEmailProcessResult, error)
	mustEmbedUnimplementedSrvUserRegEmailServer()
}

// UnimplementedSrvUserRegEmailServer must be embedded to have forward compatible implementations.
type UnimplementedSrvUserRegEmailServer struct {
}

func (UnimplementedSrvUserRegEmailServer) GetById(context.Context, *resources.UserRegEmailIdReq) (*resources.UserRegEmail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedSrvUserRegEmailServer) List(context.Context, *resources.UserRegEmailListReq) (*resources.UserRegEmailListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSrvUserRegEmailServer) Create(context.Context, *resources.UserRegEmailSaveReq) (*resources.UserRegEmail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSrvUserRegEmailServer) Update(context.Context, *resources.UserRegEmailSaveReq) (*resources.UserRegEmail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSrvUserRegEmailServer) Delete(context.Context, *resources.UserRegEmailIdReq) (*resources.UserRegEmailProcessResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSrvUserRegEmailServer) Destroy(context.Context, *resources.UserRegEmailIdsReq) (*resources.UserRegEmailProcessResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedSrvUserRegEmailServer) mustEmbedUnimplementedSrvUserRegEmailServer() {}

// UnsafeSrvUserRegEmailServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SrvUserRegEmailServer will
// result in compilation errors.
type UnsafeSrvUserRegEmailServer interface {
	mustEmbedUnimplementedSrvUserRegEmailServer()
}

func RegisterSrvUserRegEmailServer(s grpc.ServiceRegistrar, srv SrvUserRegEmailServer) {
	s.RegisterService(&SrvUserRegEmail_ServiceDesc, srv)
}

func _SrvUserRegEmail_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.UserRegEmailIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserRegEmailServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.api.user.userservicev1.SrvUserRegEmail/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserRegEmailServer).GetById(ctx, req.(*resources.UserRegEmailIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvUserRegEmail_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.UserRegEmailListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserRegEmailServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.api.user.userservicev1.SrvUserRegEmail/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserRegEmailServer).List(ctx, req.(*resources.UserRegEmailListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvUserRegEmail_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.UserRegEmailSaveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserRegEmailServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.api.user.userservicev1.SrvUserRegEmail/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserRegEmailServer).Create(ctx, req.(*resources.UserRegEmailSaveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvUserRegEmail_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.UserRegEmailSaveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserRegEmailServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.api.user.userservicev1.SrvUserRegEmail/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserRegEmailServer).Update(ctx, req.(*resources.UserRegEmailSaveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvUserRegEmail_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.UserRegEmailIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserRegEmailServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.api.user.userservicev1.SrvUserRegEmail/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserRegEmailServer).Delete(ctx, req.(*resources.UserRegEmailIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvUserRegEmail_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.UserRegEmailIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserRegEmailServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.api.user.userservicev1.SrvUserRegEmail/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserRegEmailServer).Destroy(ctx, req.(*resources.UserRegEmailIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SrvUserRegEmail_ServiceDesc is the grpc.ServiceDesc for SrvUserRegEmail service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SrvUserRegEmail_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.api.user.userservicev1.SrvUserRegEmail",
	HandlerType: (*SrvUserRegEmailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _SrvUserRegEmail_GetById_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SrvUserRegEmail_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _SrvUserRegEmail_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SrvUserRegEmail_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SrvUserRegEmail_Delete_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _SrvUserRegEmail_Destroy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/v1/services/user_reg_email.service.v1.proto",
}
