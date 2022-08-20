package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "shop/api/inventory/v1"
	"shop/app/inventory/internal/biz"
)

type InventoryService struct {
	v1.UnimplementedInventoryServer
	uc *biz.InventoryUsecase
}

func NewInventoryService(uc *biz.InventoryUsecase) *InventoryService {
	return &InventoryService{uc: uc}
}

func (i *InventoryService) SetInv(ctx context.Context, in *v1.GoodsInvInfo) (*emptypb.Empty, error) {
	err := i.uc.SetInventory(ctx, biz.Inventory{
		Goods:  in.Id,
		Stocks: in.Stocks,
	})
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (i *InventoryService) InvDetail(ctx context.Context, in *v1.GoodsInvInfo) (*v1.GoodsInvInfo, error) {
	inv, err := i.uc.GetInventory(ctx, biz.Inventory{
		Goods: in.Id,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GoodsInvInfo{
		Id:     inv.Goods,
		Stocks: inv.Stocks,
	}, nil
}

func (i *InventoryService) Sell(ctx context.Context, in *v1.SellInfo) (*emptypb.Empty, error) {
	inv := make([]biz.Inventory, 5)
	for i, x := range in.GoodsInfo {
		inv[i].Goods = x.Id
		inv[i].Stocks = x.Stocks
	}
	err := i.uc.SellInventory(ctx, inv)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (i *InventoryService) ReBack(ctx context.Context, in *v1.SellInfo) (*emptypb.Empty, error) {
	inv := make([]biz.Inventory, 5)
	for i, x := range in.GoodsInfo {
		inv[i].Goods = x.Id
		inv[i].Stocks = x.Stocks
	}
	err := i.uc.ReBackInventory(ctx, inv)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
