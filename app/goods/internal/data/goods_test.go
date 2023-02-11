package data

import (
	"context"
	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"shop/app/goods/internal/biz"
	"strconv"
	"testing"
	"time"
)

func TNewDb() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/shop_goods?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err.Error())
	}
	return db, nil
}

func TNewRedis() (*redis.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	rc := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := rc.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return rc, nil
}

func TNewGoodsRepo() (goodsRepo, error) {
	db, err := TNewDb()
	if err != nil {
		return goodsRepo{}, err
	}
	newRedis, err := TNewRedis()
	if err != nil {
		return goodsRepo{}, err
	}
	d := Data{
		db: db,
		rc: newRedis,
	}
	g := goodsRepo{
		data: &d,
		log:  nil,
	}
	return g, nil
}

func TestGetGoodsList(t *testing.T) {
	g, err := TNewGoodsRepo()
	if err != nil {
		t.Error(err)
	}
	list, err := g.GetGoodsList([]int64{50, 51, 52, 53})
	if err != nil {
		t.Error(err)
	}
	t.Log(list)
}

func TestGetGoods(t *testing.T) {
	g, err := TNewGoodsRepo()
	if err != nil {
		t.Error(err)
	}
	goods, err := g.GetGoods(50)
	if err != nil {
		t.Error(err)
	}
	t.Log(goods)
}

func TestCreateGoods(t *testing.T) {
	g, err := TNewGoodsRepo()
	if err != nil {
		t.Error(err)
	}
	err = g.CreateGoods(biz.Goods{
		Model:           gorm.Model{},
		CategoryID:      9,
		BrandsID:        1,
		OnSale:          false,
		ShipFree:        false,
		IsNew:           false,
		IsHot:           false,
		Name:            "Test",
		GoodsSn:         "1011111",
		ChickNum:        0,
		SoldNum:         0,
		FavNum:          0,
		MarketPrice:     0,
		ShopPrice:       0,
		GoodsBrief:      "",
		Images:          "/test",
		GoodsDetail:     "test",
		GoodsFrontImage: "/test",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateGoods(t *testing.T) {
	g, err := TNewGoodsRepo()
	if err != nil {
		t.Error(err)
	}
	err = g.UpdateGoods(biz.Goods{
		Model: gorm.Model{
			ID: 50,
		},
		CategoryID:      9,
		BrandsID:        1,
		OnSale:          false,
		ShipFree:        false,
		IsNew:           false,
		IsHot:           false,
		Name:            "Rog笔记本",
		GoodsSn:         "1011112",
		ChickNum:        0,
		SoldNum:         0,
		FavNum:          0,
		MarketPrice:     0,
		ShopPrice:       0,
		GoodsBrief:      "",
		Images:          "/test",
		GoodsDetail:     "test",
		GoodsFrontImage: "/test",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteGoods(t *testing.T) {
	g, err := TNewGoodsRepo()
	if err != nil {
		t.Error(err)
	}
	err = g.DeleteGoods(50)
	if err != nil {
		t.Error(err)
	}
}

func TestGetCategoryList(t *testing.T) {
	g, err := TNewGoodsRepo()
	if err != nil {
		t.Error(err)
	}
	list, err := g.GetCategoryList()
	if err != nil {
		return
	}
	t.Log(list)
}

func TestGetCategory(t *testing.T) {
	g, err := TNewGoodsRepo()
	if err != nil {
		t.Error(err)
	}
	category, err := g.GetCategory(9)
	if err != nil {
		return
	}
	t.Log(category)
}

func TestUpdateCategory(t *testing.T) {
	g, err := TNewGoodsRepo()
	if err != nil {
		t.Error(err)
	}
	err = g.UpdateCategory(biz.Category{
		Model: gorm.Model{
			ID: 10,
		},
		Name:  "家电",
		Level: 2,
		IsTab: true,
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteCategory(t *testing.T) {
	g, err := TNewGoodsRepo()
	if err != nil {
		t.Error(err)
	}
	err = g.DeleteCategory(9)
	if err != nil {
		t.Error(err)
	}
}

func TestGetBrandList(t *testing.T) {
	g, err := TNewGoodsRepo()
	if err != nil {
		t.Error(err)
	}
	list, err := g.GetBrandList()
	if err != nil {
		t.Error(err)
	}
	t.Log(list)
}

func TestGetBrand(t *testing.T) {
	g, err := TNewGoodsRepo()
	if err != nil {
		t.Error(err)
	}
	brand, err := g.GetBrand(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(brand)
}

func TestCreateBrand(t *testing.T) {
	g, err := TNewGoodsRepo()
	if err != nil {
		t.Error(err)
	}
	err = g.CreateBrand(biz.Brands{
		Name: "英伟达",
		Logo: "/test",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateBrand(t *testing.T) {
	g, err := TNewGoodsRepo()
	if err != nil {
		t.Error(err)
	}
	err = g.UpdateBrand(biz.Brands{
		Model: gorm.Model{ID: 1},
		Name:  "jindongf",
		Logo:  "/test",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteBrand(t *testing.T) {
	g, err := TNewGoodsRepo()
	if err != nil {
		t.Error(err)
	}
	err = g.DeleteBrand(1)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateGoodsCache(t *testing.T) {
	g, err := TNewGoodsRepo()
	if err != nil {
		t.Error(err)
	}
	goods, err := g.GetGoods(50)
	if err != nil {
		t.Error(err)
	}
	t.Log(goods)
	err = g.CreateGoodsCache(goods)
	if err != nil {
		t.Error("set", err)
	}
}

func TestGetGoodsCache(t *testing.T) {
	g, err := TNewGoodsRepo()
	if err != nil {
		t.Error(err)
	}
	cache, err := g.GetGoodsCache(strconv.Itoa(50))
	if err != nil {
		t.Error("get", err)
	}
	t.Log(cache)
}
