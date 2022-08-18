package service

import (
	"context"
	"fmt"
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

func (i *InventoryService) SetInv(context.Context, *v1.GoodsInvInfo) (*emptypb.Empty, error) {
	return nil, nil
}

func (i *InventoryService) InvDetail(ctx context.Context, in *v1.GoodsInvInfo) (*v1.GoodsInvInfo, error) {
	inv, err := i.uc.InvDetail(ctx, biz.Inventory{
		Goods: in.Id,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(inv)
	return &v1.GoodsInvInfo{
		Id:     inv.Goods,
		Stocks: inv.Stocks,
	}, nil
}

func (i *InventoryService) Sell(context.Context, *v1.SellInfo) (*emptypb.Empty, error) {
	return nil, nil
}

func (i *InventoryService) ReBack(context.Context, *v1.SellInfo) (*emptypb.Empty, error) {
	return nil, nil
}
