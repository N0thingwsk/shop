// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.1
// - protoc             v3.21.2
// source: inventory/v1/inventory.proto

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

const OperationInventoryInvDetail = "/api.inventory.v1.Inventory/InvDetail"

type InventoryHTTPServer interface {
	InvDetail(context.Context, *GoodsInvInfo) (*GoodsInvInfo, error)
}

func RegisterInventoryHTTPServer(s *http.Server, srv InventoryHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/inventory/getinventory", _Inventory_InvDetail0_HTTP_Handler(srv))
}

func _Inventory_InvDetail0_HTTP_Handler(srv InventoryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GoodsInvInfo
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationInventoryInvDetail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.InvDetail(ctx, req.(*GoodsInvInfo))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GoodsInvInfo)
		return ctx.Result(200, reply)
	}
}

type InventoryHTTPClient interface {
	InvDetail(ctx context.Context, req *GoodsInvInfo, opts ...http.CallOption) (rsp *GoodsInvInfo, err error)
}

type InventoryHTTPClientImpl struct {
	cc *http.Client
}

func NewInventoryHTTPClient(client *http.Client) InventoryHTTPClient {
	return &InventoryHTTPClientImpl{client}
}

func (c *InventoryHTTPClientImpl) InvDetail(ctx context.Context, in *GoodsInvInfo, opts ...http.CallOption) (*GoodsInvInfo, error) {
	var out GoodsInvInfo
	pattern := "/v1/inventory/getinventory"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationInventoryInvDetail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}