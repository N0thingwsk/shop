package test

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"shop/app/goods/internal/biz"
	"testing"
)

func TestGorm(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/shop_goods?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err.Error())
	}
	//err = db.AutoMigrate(&biz.Category{}, biz.Brands{}, &biz.GoodsCategoryBrand{}, &biz.Goods{})
	//if err != nil {
	//	panic(err.Error())
	//}
	tx := db.Begin()
	var cat biz.Category
	tx.First(&cat, 9)
	var ban biz.Brands
	tx.First(&ban, 1)
	t.Log(cat, ban)
	for i := 1; i <= 10; i++ {
		tx.Create(&biz.Goods{
			Category:        cat,
			Brands:          ban,
			OnSale:          false,
			ShipFree:        false,
			IsNew:           false,
			IsHot:           false,
			Name:            fmt.Sprintf("test%d", i),
			GoodsSn:         fmt.Sprintf("102030123%d", i),
			ChickNum:        int64(i),
			SoldNum:         10,
			FavNum:          10,
			MarketPrice:     10,
			ShopPrice:       10,
			GoodsBrief:      "test",
			Images:          "/test",
			GoodsDetail:     "test",
			GoodsFrontImage: "/test",
		})
	}
	tx.Commit()

}
