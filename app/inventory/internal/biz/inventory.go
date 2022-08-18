package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	v1 "shop/api/user/v1"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
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
}

type InventoryUsecase struct {
	repo InventoryRepo
	log  *log.Helper
}

func NewInventoryUsecase(repo InventoryRepo, logger log.Logger) *InventoryUsecase {
	return &InventoryUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (ic *InventoryUsecase) CreateGreeter(ctx context.Context, g *Inventory) (*Inventory, error) {
	return nil, nil
}
