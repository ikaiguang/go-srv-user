// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package userservicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	resources "github.com/ikaiguang/go-srv-user/api/user/v1/resources"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSrvUserAuthLoginOrRegByPhone = "/user.api.user.userservicev1.SrvUserAuth/LoginOrRegByPhone"

type SrvUserAuthHTTPServer interface {
	LoginOrRegByPhone(context.Context, *resources.LoginOrRegByPhoneReq) (*resources.LoginResp, error)
}

func RegisterSrvUserAuthHTTPServer(s *http.Server, srv SrvUserAuthHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/user/login-or-reg-phone", _SrvUserAuth_LoginOrRegByPhone0_HTTP_Handler(srv))
}

func _SrvUserAuth_LoginOrRegByPhone0_HTTP_Handler(srv SrvUserAuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.LoginOrRegByPhoneReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvUserAuthLoginOrRegByPhone)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginOrRegByPhone(ctx, req.(*resources.LoginOrRegByPhoneReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.LoginResp)
		return ctx.Result(200, reply)
	}
}

type SrvUserAuthHTTPClient interface {
	LoginOrRegByPhone(ctx context.Context, req *resources.LoginOrRegByPhoneReq, opts ...http.CallOption) (rsp *resources.LoginResp, err error)
}

type SrvUserAuthHTTPClientImpl struct {
	cc *http.Client
}

func NewSrvUserAuthHTTPClient(client *http.Client) SrvUserAuthHTTPClient {
	return &SrvUserAuthHTTPClientImpl{client}
}

func (c *SrvUserAuthHTTPClientImpl) LoginOrRegByPhone(ctx context.Context, in *resources.LoginOrRegByPhoneReq, opts ...http.CallOption) (*resources.LoginResp, error) {
	var out resources.LoginResp
	pattern := "/api/v1/user/login-or-reg-phone"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSrvUserAuthLoginOrRegByPhone))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}