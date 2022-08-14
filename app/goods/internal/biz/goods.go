package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	v1 "shop/api/user/v1"
)

var (
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type Category struct {
	gorm.Model
	Name             string `gorm:"type:varchar(20);not null"`
	ParentCategoryID int32
	ParentCategory   *Category
	Level            int32 `gorm:"type:int;not null;default:1"`
	IsTab            bool  `gorm:"default:false;not null"`
}

type Brands struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Logo string `gorm:"type:varchar(200);default:'';not null"`
}

type GoodsCategoryBrand struct {
	gorm.Model
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands     Brands
}

type Banner struct {
	gorm.Model
	Image string `gorm:"type:varchar(200);not null"`
	Url   string `gorm:"type:varchar(200);not null"`
	Index int32  `gorm:"type:int;default:1;not null"`
}

type Goods struct {
	gorm.Model
	CategoryID      int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category        Category
	BrandsID        int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands          Brands
	OnSale          bool    `gorm:"default:false;not null"`
	ShipFree        bool    `gorm:"default;false;not null"`
	IsNew           bool    `gorm:"default;false;not null"`
	IsHot           bool    `gorm:"default;false;not null"`
	Name            string  `gorm:"type:varchar(50);not null"`
	GoodsSn         string  `gorm:"type:varchar(50);not null"`
	ChickNum        int32   `gorm:"type:int;default:0;not null"`
	SoldNum         int32   `gorm:"type:int;default:0;not null"`
	FavNum          int32   `gorm:"type:int;default:0;not null"`
	MarketPrice     float32 `gorm:"not null"`
	ShopPrice       float32 `gorm:"not null"`
	GoodsBrief      string  `gorm:"type:varchar(50);not null"`
	Images          []string
	DescImages      []string
	GoodsFrontImage string `gorm:"type:varchar(200);not null"`
}

type GoodsRepo interface {
}

type GoodsUsecase struct {
	repo GoodsRepo
	log  *log.Helper
}

func NewGreeterUsecase(repo GoodsRepo, logger log.Logger) *GoodsUsecase {
	return &GoodsUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GoodsUsecase) CreateGreeter(ctx context.Context, g *Goods) (*Goods, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v")
	return g, nil
}
