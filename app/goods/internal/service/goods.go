package service

import (
	"context"
	"shop/app/goods/internal/biz"

	pb "shop/api/goods/v1"
)

type GoodsService struct {
	pb.UnimplementedGoodsServer
	uc *biz.GoodsUsecase
}

func NewGoodsService(uc *biz.GoodsUsecase) *GoodsService {
	return &GoodsService{uc: uc}
}

func (s *GoodsService) GetGoodsList(ctx context.Context, req *pb.GetGoodsListRequest) (*pb.GetGoodsListReply, error) {
	var gs []*pb.GoodsValue
	list, err := s.uc.GetGoodsList(req.Id)
	if err != nil {
		return nil, err
	}
	for _, x := range list {
		gs = append(gs, &pb.GoodsValue{
			OnSale:          x.OnSale,
			ShipFree:        x.ShipFree,
			IsNew:           x.IsNew,
			IsHot:           x.IsHot,
			Name:            x.Name,
			GoodsSn:         x.GoodsSn,
			ChickNum:        x.ChickNum,
			SoldNum:         x.SoldNum,
			FavNum:          x.FavNum,
			MarketPrice:     x.MarketPrice,
			ShopPrice:       x.ShopPrice,
			GoodsBrief:      x.GoodsBrief,
			Images:          x.Images,
			GoodsDetail:     x.GoodsDetail,
			GoodsFrontImage: x.GoodsFrontImage,
			Id:              int64(x.ID),
		})
	}
	return &pb.GetGoodsListReply{
		Code:    200,
		Message: "获取成功",
		Goods:   nil,
	}, nil
}
func (s *GoodsService) GetGoods(ctx context.Context, req *pb.GoodsRequest) (*pb.GoodsReply, error) {

	return &pb.GoodsReply{}, nil
}
func (s *GoodsService) CreateGoods(ctx context.Context, req *pb.CreateGoodsRequest) (*pb.CreateGoodsReply, error) {
	return &pb.CreateGoodsReply{}, nil
}
func (s *GoodsService) UpdateGoods(ctx context.Context, req *pb.UpdateGoodsRequest) (*pb.UpdateGoodsReply, error) {
	return &pb.UpdateGoodsReply{}, nil
}
func (s *GoodsService) DeleteGoods(ctx context.Context, req *pb.DeleteGoodsRequest) (*pb.DeleteGoodsReply, error) {
	return &pb.DeleteGoodsReply{}, nil
}
func (s *GoodsService) GetCategoryList(ctx context.Context, req *pb.GetCategoryListRequest) (*pb.GetCategoryListReply, error) {
	return &pb.GetCategoryListReply{}, nil
}
func (s *GoodsService) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.GetCategoryReply, error) {
	return &pb.GetCategoryReply{}, nil
}
func (s *GoodsService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryReply, error) {
	return &pb.CreateCategoryReply{}, nil
}
func (s *GoodsService) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryReply, error) {
	return &pb.DeleteCategoryReply{}, nil
}
func (s *GoodsService) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryReply, error) {
	return &pb.UpdateCategoryReply{}, nil
}
func (s *GoodsService) GetBrandList(ctx context.Context, req *pb.GetBrandListRequest) (*pb.GetBrandListReply, error) {
	return &pb.GetBrandListReply{}, nil
}
func (s *GoodsService) GetBrand(ctx context.Context, req *pb.GetBrandRequest) (*pb.GetBrandReply, error) {
	return &pb.GetBrandReply{}, nil
}
func (s *GoodsService) CreateBrand(ctx context.Context, req *pb.CreateBrandRequest) (*pb.CreateBrandReply, error) {
	return &pb.CreateBrandReply{}, nil
}
func (s *GoodsService) UpdateBrand(ctx context.Context, req *pb.UpdateBrandRequest) (*pb.UpdateBrandReply, error) {
	return &pb.UpdateBrandReply{}, nil
}
func (s *GoodsService) DeleteBrand(ctx context.Context, req *pb.DeleteBrandRequest) (*pb.DeleteBrandReply, error) {
	return &pb.DeleteBrandReply{}, nil
}
