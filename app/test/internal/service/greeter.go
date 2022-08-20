package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	kg "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	v1 "shop/api/helloworld/v1"
	inv "shop/api/inventory/v1"
	"shop/app/test/internal/biz"
	"shop/app/test/internal/conf"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer
	uc        *biz.GreeterUsecase
	inventory inv.InventoryClient
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, conn *grpc.ClientConn) *GreeterService {
	inventory := inv.NewInventoryClient(conn)
	return &GreeterService{uc: uc, inventory: inventory}
}

func NewGrpcClient(c *conf.Data) *grpc.ClientConn {
	consulCli, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	r := consul.New(consulCli)
	conn, err := kg.DialInsecure(
		context.Background(),
		kg.WithEndpoint("discovery:///inventory"),
		kg.WithDiscovery(r),
	)
	if err != nil {
		panic(err.Error())
	}
	return conn
}

func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	num, err := s.inventory.InvDetail(ctx, &inv.GoodsInvInfo{
		Id: 1,
	})
	if err != nil {
		return &v1.HelloReply{Message: "错误"}, err
	}
	fmt.Println(num)
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
