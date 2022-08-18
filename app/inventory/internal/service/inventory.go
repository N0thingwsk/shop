package service

import (
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
