package data

import (
	"context"
	"errors"
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

func (i *inventoryRepo) GetInventory(ctx context.Context, inv biz.Inventory) (biz.Inventory, error) {
	result := i.data.db.Where("goods = ?", inv.Goods).First(&inv)
	if result.Error != nil {
		return biz.Inventory{}, result.Error
	} else if result.RowsAffected == 0 {
		return biz.Inventory{}, errors.New("未查询到库存")
	}
	return inv, nil
}

func (i *inventoryRepo) CreateInventory(context.Context, biz.Inventory) error {
	return nil
}

func (i *inventoryRepo) UpdateInventory(ctx context.Context, inv biz.Inventory) error {
	result := i.data.db.Save(&inv)
	if result.Error != nil {
		return errors.New("更新库存失败")
	}
	return nil
}

func (i *inventoryRepo) DeleteInventory(context.Context, biz.Inventory) error {
	return nil
}

func (i *inventoryRepo) SellInventory(ctx context.Context, inv []biz.Inventory) error {
	tx := i.data.db.Begin()
	for _, x := range inv {
		result, err := i.GetInventory(ctx, x)
		if err != nil {
			tx.Rollback()
			return err
		}
		if result.Stocks < x.Stocks {
			tx.Rollback()
			return errors.New("库存不足")
		}
		result.Stocks -= x.Stocks
		tx.Save(&result)
	}
	tx.Commit()
	return nil
}

func (i *inventoryRepo) ReBackInventory(ctx context.Context, inv []biz.Inventory) error {
	tx := i.data.db.Begin()
	for _, x := range inv {
		result, err := i.GetInventory(ctx, x)
		if err != nil {
			tx.Rollback()
			return err
		}
		result.Stocks += x.Stocks
		tx.Save(&result)
	}
	tx.Commit()
	return nil
}
