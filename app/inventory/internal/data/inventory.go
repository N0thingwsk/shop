package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"shop/app/inventory/internal/biz"
)

type inventoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewInventoryRepo(data *Data, logger log.Logger) biz.InventoryRepo {
	return &inventoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
