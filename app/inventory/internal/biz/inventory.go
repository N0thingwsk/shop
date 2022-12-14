package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	Goods   int32 `gorm:"type:int;"`
	Stocks  int32 `gorm:"type:int"`
	Version int32 `gorm:"type:int"`
}

type InventoryHistory struct {
	user   int32
	goods  int32
	nums   int32
	order  int32
	status int32
}

type InventoryRepo interface {
	GetInventory(context.Context, Inventory) (Inventory, error)
	CreateInventory(context.Context, Inventory) error
	UpdateInventory(context.Context, Inventory) error
	DeleteInventory(context.Context, Inventory) error
	SellInventory(context.Context, []Inventory) error
	ReBackInventory(context.Context, []Inventory) error
}

type InventoryUsecase struct {
	repo InventoryRepo
	log  *log.Helper
}

func NewInventoryUsecase(repo InventoryRepo, logger log.Logger) *InventoryUsecase {
	return &InventoryUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (i *InventoryUsecase) GetInventory(ctx context.Context, inv Inventory) (Inventory, error) {
	result, err := i.repo.GetInventory(ctx, inv)
	if err != nil {
		return Inventory{}, err
	}
	return result, nil
}

func (i *InventoryUsecase) CreateInventory(ctx context.Context, inv Inventory) (Inventory, error) {
	err := i.repo.CreateInventory(ctx, inv)
	if err != nil {
		return Inventory{}, err
	}
	return inv, nil
}

func (i *InventoryUsecase) SetInventory(ctx context.Context, inv Inventory) error {
	_, err := i.repo.GetInventory(ctx, inv)
	if err != nil {
		return err
	}
	err = i.repo.UpdateInventory(ctx, inv)
	if err != nil {
		return err
	}
	return nil
}

func (i *InventoryUsecase) SellInventory(ctx context.Context, inv []Inventory) error {
	err := i.repo.SellInventory(ctx, inv)
	if err != nil {
		return err
	}
	return nil
}

func (i *InventoryUsecase) ReBackInventory(ctx context.Context, inv []Inventory) error {
	err := i.repo.ReBackInventory(ctx, inv)
	if err != nil {
		return err
	}
	return nil
}
