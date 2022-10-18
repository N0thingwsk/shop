package data

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"shop/app/goods/internal/biz"
	"testing"
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

func TestGetGoodsList(t *testing.T) {
	db, _ := TNewDb()
	//d := Data{
	//	db: db,
	//}
	//var l log.Logger
	//g := NewGoodsRepo(&d, l)
	//list, err := g.GetGoodsList([]int64{41, 42, 43, 44})
	//if err != nil {
	//	t.Error("error")
	//}
	//t.Log(list)
	var gds biz.Goods
	db.First(&gds, 78)
	t.Log(gds)
	//id := []int64{88, 89, 90}
	//var goods []biz.Goods
	//for _, x := range id {
	//	var gds biz.Goods
	//	result := db.First(&gds, x)
	//	if result.Error != nil {
	//		t.Error("error")
	//		continue
	//	}
	//	t.Log(gds)
	//	goods = append(goods, gds)
	//}
	//marshal, err := json.Marshal(goods)
	//if err != nil {
	//	return
	//}
	//t.Log(string(marshal))
}

type User struct {
	gorm.Model
	Name string
	Pro  []Profile `gorm:"foreignkey:id"`
}

type Profile struct {
	gorm.Model
	Pro string
}

func TestSC(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err.Error())
	}
	//db.AutoMigrate(&User{}, &Profile{})
	db.Create(&User{
		Name: "Alisi",
		Pro: []Profile{
			{
				Pro: "test",
			},
		},
	})
	var ali User
	fmt.Println(ali)
}
