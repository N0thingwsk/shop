package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

// Category 商品分类
type Category struct {
	gorm.Model
	//分类名称
	Name string `gorm:"type:varchar(20);not null"`
	//分类级别
	Level int64 `gorm:"type:int;not null;default:1"`
	//是否显示
	IsTab bool `gorm:"default:false;not null"`
}

// Brands 商品品牌
type Brands struct {
	gorm.Model
	//商品品牌名称
	Name string `gorm:"type:varchar(20);not null"`
	//商品品牌Logo
	Logo string `gorm:"type:varchar(200);default:'';not null"`
}

// GoodsCategoryBrand 商品品牌分类中间表
type GoodsCategoryBrand struct {
	gorm.Model
	CategoryID int64 `gorm:"type:int;index:idx_category_brand,unique"`
	Category   Category
	BrandsID   int64 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands     Brands
}

// Goods 商品信息
type Goods struct {
	gorm.Model
	CategoryID int64 `gorm:"type:int;not null"`
	Category   *Category
	BrandsID   int64 `gorm:"type:int;not null"`
	Brands     *Brands
	//是否上架
	OnSale bool `gorm:"default:false;not null"`
	//是否免运费
	ShipFree bool `gorm:"default:false;not null"`
	//是否为新品商品
	IsNew bool `gorm:"default:false;not null"`
	//是否为热卖商品
	IsHot bool `gorm:"default:false;not null"`
	//商品名
	Name string `gorm:"type:varchar(50);not null"`
	//商品编号
	GoodsSn string `gorm:"type:varchar(50);not null"`
	//点击量
	ChickNum int64 `gorm:"type:int;default:0;not null"`
	//销量
	SoldNum int64 `gorm:"type:int;default:0;not null"`
	//收藏数量
	FavNum int64 `gorm:"type:int;default:0;not null"`
	//商品价格
	MarketPrice float32 `gorm:"not null"`
	//店铺价格
	ShopPrice float32 `gorm:"not null"`
	//商品简介
	GoodsBrief string `gorm:"type:varchar(100);not null"`
	//商品图片
	Images string `gorm:"type:varchar(1000);not null"`
	//商品详情
	GoodsDetail string `gorm:"type:varchar(1000);not null"`
	//商品封面
	GoodsFrontImage string `gorm:"type:varchar(200);not null"`
}

type GoodsRepo interface {
	GetGoodsList([]int64) ([]Goods, error)
	GetGoods(int64) (Goods, error)
	CreateGoods(Goods) error
	UpdateGoods(Goods) error
	DeleteGoods(int64) error
	GetCategoryList() ([]Category, error)
	GetCategory(int64) (Category, error)
	UpdateCategory(Category) error
	DeleteCategory(int64) error
	GetBrandList() ([]Brands, error)
	GetBrand(int64) (Brands, error)
	CreateBrand(Brands) error
	UpdateBrand(Brands) error
	DeleteBrand(int64) error
	GetGoodsCache(key string) (string, error)
	GetGoodsListCache(key []int64) (string, error)
	CreateGoodsCache(goods Goods) error
	DeleteGoodsCache(key string) error
}

type GoodsUsecase struct {
	repo GoodsRepo
	log  *log.Helper
}

func NewGoodsUsecase(repo GoodsRepo, logger log.Logger) *GoodsUsecase {
	return &GoodsUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (g *GoodsUsecase) GetGoodsList(id []int64) ([]Goods, error) {
	return []Goods{
		{
			Model:           gorm.Model{},
			CategoryID:      0,
			Category:        Category{},
			BrandsID:        0,
			Brands:          Brands{},
			OnSale:          false,
			ShipFree:        false,
			IsNew:           false,
			IsHot:           false,
			Name:            "",
			GoodsSn:         "",
			ChickNum:        0,
			SoldNum:         0,
			FavNum:          0,
			MarketPrice:     0,
			ShopPrice:       0,
			GoodsBrief:      "test",
			Images:          "test",
			GoodsDetail:     "test",
			GoodsFrontImage: "test",
		},
	}, nil
}
