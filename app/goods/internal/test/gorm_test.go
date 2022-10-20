package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"testing"
	"time"
)

type User struct {
	gorm.Model
	Name string
}

func TestGorm(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
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
	user := User{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "LLLLL",
	}

	db.Model(&User{}).Updates(&user)
	fmt.Println(user)

}

func TestUser(t *testing.T) {
	r := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r.HSet(ctx, "user",
		"name", "",
		"age", "12",
	)
	result, err := r.HGetAll(ctx, "user").Result()
	if err != nil {
		return
	}
	fmt.Println(result)
}
