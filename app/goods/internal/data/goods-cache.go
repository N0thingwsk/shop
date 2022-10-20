package data

//func (g *goodsRepo) GetGoodsCache(key string) (map[string]string, error) {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
//	defer cancel()
//	result, err := g.data.rc.HGetAll(ctx, key).Result()
//	if err != nil {
//		return result, err
//	}
//	return result, nil
//}
//
//func (g *goodsRepo) SetGoodsCache(goods biz.Goods) error {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
//	defer cancel()
//	err := g.data.rc.HSet(ctx,
//		strconv.Itoa(int(goods.ID)),
//		"Category", goods.Category.Name,
//		"Brands ", goods.Brands.Name,
//		"OnSale", goods.OnSale,
//		"ShipFree", goods.ShipFree,
//		"IsNew", goods.IsNew,
//		"IsHot", goods.IsHot,
//		"Name", goods.Name,
//		"GoodsSn", goods.GoodsSn,
//		"ChickNum", goods.ChickNum,
//		"SoldNum", goods.SoldNum,
//		"FavNum", goods.FavNum,
//		"MarketPrice", goods.MarketPrice,
//		"ShopPrice", goods.ShopPrice,
//		"GoodsBrief", goods.GoodsBrief,
//		"Images", goods.Images,
//		"GoodsDetail", goods.GoodsDetail,
//		"GoodsFrontImage", goods.GoodsFrontImage,
//	).Err()
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (g *goodsRepo) DeleteGoodsCache(key string) error {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
//	defer cancel()
//	err := g.data.rc.Del(ctx, key).Err()
//	if err != nil {
//		return err
//	}
//	return nil
//}
