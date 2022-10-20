// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.1
// - protoc             v3.21.7
// source: userop/v1/userop.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUseropUsdropUser = "/api.userop.v1.Userop/UsdropUser"

type UseropHTTPServer interface {
	UsdropUser(context.Context, *CreateUseropReply) (*CreateUseropReply, error)
}

func RegisterUseropHTTPServer(s *http.Server, srv UseropHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/user/getuser", _Userop_UsdropUser0_HTTP_Handler(srv))
}

func _Userop_UsdropUser0_HTTP_Handler(srv UseropHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUseropReply
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUseropUsdropUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UsdropUser(ctx, req.(*CreateUseropReply))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUseropReply)
		return ctx.Result(200, reply)
	}
}

type UseropHTTPClient interface {
	UsdropUser(ctx context.Context, req *CreateUseropReply, opts ...http.CallOption) (rsp *CreateUseropReply, err error)
}

type UseropHTTPClientImpl struct {
	cc *http.Client
}

func NewUseropHTTPClient(client *http.Client) UseropHTTPClient {
	return &UseropHTTPClientImpl{client}
}

func (c *UseropHTTPClientImpl) UsdropUser(ctx context.Context, in *CreateUseropReply, opts ...http.CallOption) (*CreateUseropReply, error) {
	var out CreateUseropReply
	pattern := "/v1/user/getuser"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUseropUsdropUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}