// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/user/v1/services/user_auth.service.v1.proto

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

// SrvUserAuthClient is the client API for SrvUserAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SrvUserAuthClient interface {
	// LoginOrRegByPhone 手机登陆(自动注册)
	LoginOrRegByPhone(ctx context.Context, in *resources.LoginOrRegByPhoneReq, opts ...grpc.CallOption) (*resources.LoginResp, error)
}

type srvUserAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewSrvUserAuthClient(cc grpc.ClientConnInterface) SrvUserAuthClient {
	return &srvUserAuthClient{cc}
}

func (c *srvUserAuthClient) LoginOrRegByPhone(ctx context.Context, in *resources.LoginOrRegByPhoneReq, opts ...grpc.CallOption) (*resources.LoginResp, error) {
	out := new(resources.LoginResp)
	err := c.cc.Invoke(ctx, "/user.api.user.userservicev1.SrvUserAuth/LoginOrRegByPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SrvUserAuthServer is the server API for SrvUserAuth service.
// All implementations must embed UnimplementedSrvUserAuthServer
// for forward compatibility
type SrvUserAuthServer interface {
	// LoginOrRegByPhone 手机登陆(自动注册)
	LoginOrRegByPhone(context.Context, *resources.LoginOrRegByPhoneReq) (*resources.LoginResp, error)
	mustEmbedUnimplementedSrvUserAuthServer()
}

// UnimplementedSrvUserAuthServer must be embedded to have forward compatible implementations.
type UnimplementedSrvUserAuthServer struct {
}

func (UnimplementedSrvUserAuthServer) LoginOrRegByPhone(context.Context, *resources.LoginOrRegByPhoneReq) (*resources.LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginOrRegByPhone not implemented")
}
func (UnimplementedSrvUserAuthServer) mustEmbedUnimplementedSrvUserAuthServer() {}

// UnsafeSrvUserAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SrvUserAuthServer will
// result in compilation errors.
type UnsafeSrvUserAuthServer interface {
	mustEmbedUnimplementedSrvUserAuthServer()
}

func RegisterSrvUserAuthServer(s grpc.ServiceRegistrar, srv SrvUserAuthServer) {
	s.RegisterService(&SrvUserAuth_ServiceDesc, srv)
}

func _SrvUserAuth_LoginOrRegByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.LoginOrRegByPhoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserAuthServer).LoginOrRegByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.api.user.userservicev1.SrvUserAuth/LoginOrRegByPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserAuthServer).LoginOrRegByPhone(ctx, req.(*resources.LoginOrRegByPhoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SrvUserAuth_ServiceDesc is the grpc.ServiceDesc for SrvUserAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SrvUserAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.api.user.userservicev1.SrvUserAuth",
	HandlerType: (*SrvUserAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginOrRegByPhone",
			Handler:    _SrvUserAuth_LoginOrRegByPhone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/v1/services/user_auth.service.v1.proto",
}