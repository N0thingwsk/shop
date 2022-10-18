package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v9"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"shop/app/goods/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGoodsRepo, NewDb, NewCache)

// Data .
type Data struct {
	db *gorm.DB
	rc *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, rc *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db, rc: rc}, cleanup, nil
}

func NewDb(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err.Error())
	}
	return db
}

func NewCache(c *conf.Data) *redis.Client {
	rc := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: "",
		DB:       0,
	})
	return rc
}
