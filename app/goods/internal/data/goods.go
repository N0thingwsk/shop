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

func (g *goodsRepo) GetGoods(id int64) (biz.Goods, error) {
	var goods biz.Goods
	result := g.data.db.First(&goods, id)
	if result.Error != nil {
		return biz.Goods{}, result.Error
	}
	return goods, nil
}

func (g *goodsRepo) CreateGoods(goods biz.Goods) error {
	result := g.data.db.Create(&goods)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (g *goodsRepo) UpdateGoods(goods biz.Goods) error {
	result := g.data.db.Updates(&goods)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (g *goodsRepo) DeleteGoods(id int64) error {
	result := g.data.db.Where("id = ?", id).Delete(&biz.Goods{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (g *goodsRepo) GetCategoryList() ([]biz.Category, error) {
	var category []biz.Category
	result := g.data.db.Find(&category)
	if result.Error != nil {
		return []biz.Category{}, result.Error
	}
	return category, nil
}

func (g *goodsRepo) GetCategory(id int64) (biz.Category, error) {
	var category biz.Category
	result := g.data.db.First(&category, id)
	if result.Error != nil {
		return biz.Category{}, result.Error
	}
	return category, nil
}

func (g *goodsRepo) UpdateCategory(category biz.Category) error {
	result := g.data.db.Updates(&category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (g *goodsRepo) DeleteCategory(id int64) error {
	result := g.data.db.Where("id = ?", id).Delete(&biz.Category{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (g *goodsRepo) GetBrandList() ([]biz.Brands, error) {
	var brands []biz.Brands
	result := g.data.db.Find(&brands)
	if result.Error != nil {
		return []biz.Brands{}, result.Error
	}
	return brands, nil
}

func (g *goodsRepo) GetBrand(id int64) (biz.Brands, error) {
	var brands biz.Brands
	result := g.data.db.First(&brands, id)
	if result.Error != nil {
		return biz.Brands{}, result.Error
	}
	return brands, nil
}

func (g *goodsRepo) CreateBrand(brands biz.Brands) error {
	result := g.data.db.Create(&brands)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (g *goodsRepo) UpdateBrand(brands biz.Brands) error {
	result := g.data.db.Updates(&brands)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (g goodsRepo) DeleteBrand(id int64) error {
	result := g.data.db.Where("id = ?", id).Delete(&biz.Brands{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
