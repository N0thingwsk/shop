package biz

import (
	"gorm.io/gorm"
	"time"
)

type ShoppingCart struct {
	gorm.Model
	User    int32 `gorm:"type:int;index"`
	Goods   int32 `gorm:"type:int;index"`
	Nums    int32 `gorm:"type:int"`
	Checked bool
}

type OrderInfo struct {
	gorm.Model
	User         int32  `gorm:"type:int;index"`
	OrderSn      string `gorm:"type:varchar(30);index"`
	PlayType     string `gorm:"type:varchar(20)"`
	Status       string `gorm:"varchar(20)"`
	TradeNo      string `gorm:"varchar(100)"`
	OrderMount   float32
	PayTime      time.Time
	Address      string `gorm:"type:varchar(100)"`
	SignerName   string `gorm:"type:varchar(20)"`
	SingerMobile string `gorm:"type:varchar(11)"`
	Post         string `gorm:"type:varchar(20)"`
}

type OrderGoods struct {
	gorm.Model
	Order      int32  `gorm:"type:int;index"`
	Goods      int32  `gorm:"type:int;index"`
	GoodsName  string `gorm:"type:varchar(100);index"`
	GoodsImage string `gorm:"type:varchar(200)"`
	GoodsPrice float32
	Nums       int32 `gorm:"type:int"`
}

func (ShoppingCart) TableName() string {
	return "shoppingcart"
}

func (OrderInfo) TableName() string {
	return "orderInfo"
}

func (OrderGoods) TableName() string {
	return "ordergoods"
}
