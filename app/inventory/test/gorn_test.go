package test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"shop/app/inventory/internal/biz"
	"testing"
)

func TestGrom(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/shop_inventory?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&biz.Inventory{})
}
