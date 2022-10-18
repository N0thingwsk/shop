package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"shop/app/goods/internal/biz"
)

type goodsRepo struct {
	data *Data
	log  *log.Helper
}

// NewGoodsRepo .
func NewGoodsRepo(data *Data, logger log.Logger) biz.GoodsRepo {
	return &goodsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (g *goodsRepo) GetGoodsList(id []int64) ([]biz.Goods, error) {
	var goods []biz.Goods
	for _, x := range id {
		var gds biz.Goods
		result := g.data.db.First(&gds, x)
		goods = append(goods, gds)
		if result.Error != nil {
			return []biz.Goods{}, result.Error
		}
	}
	return goods, nil
}

func (g *goodsRepo) GetGoods(id int) (biz.Goods, error) {
	var goods biz.Goods
	result := g.data.db.First(&goods, id)
	if result.Error != nil {
		return biz.Goods{}, result.Error
	}
	return goods, nil
}

func (g *goodsRepo) CreateGoods(goods biz.Goods) error {
	g.data.db.Create(&goods)
	return nil
}

func (g *goodsRepo) UpdateGoods(goods biz.Goods) error {
	return nil
}

func (g *goodsRepo) DeleteGoods(id int) error {
	g.data.db.Where("id = ?", id).Delete(&biz.Goods{})
	return nil
}

func (g *goodsRepo) GetCategoryList(id []int) ([]biz.Category, error) {
	var goods []biz.Category
	for _, x := range id {
		var gds biz.Category
		result := g.data.db.First(&gds, x)
		goods = append(goods, gds)
		if result.Error != nil {
			return []biz.Category{}, result.Error
		}
	}
	return goods, nil
}

func (g *goodsRepo) GetCategory(id int) (biz.Category, error) {
	var category biz.Category
	result := g.data.db.First(&category, id)
	if result.Error != nil {
		return biz.Category{}, result.Error
	}
	return category, nil
}

func (g *goodsRepo) DeleteCategory(id int) error {
	return nil
}
