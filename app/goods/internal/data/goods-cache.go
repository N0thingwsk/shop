package data

import (
	"context"
	"encoding/json"
	"math/rand"
	"shop/app/goods/internal/biz"
	"strconv"
	"time"
)

func (g *goodsRepo) GetGoodsCache(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	result, err := g.data.rc.Get(ctx, key).Result()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (g *goodsRepo) GetGoodsListCache(key []int64) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	goods := make(map[string]string)
	for _, x := range key {
		result, err := g.data.rc.Get(ctx, string(x)).Result()
		if err != nil {
			return "", err
		}
		goods[string(x)] = result
	}
	marshal, err := json.Marshal(&goods)
	if err != nil {
		return "", err
	}
	return string(marshal), err
}

func (g *goodsRepo) CreateGoodsCache(goods biz.Goods) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	marshal, err := json.Marshal(&goods)
	if err != nil {
		return err
	}
	err = g.data.rc.Set(ctx, strconv.Itoa(int(goods.ID)), marshal, time.Second*time.Duration(1000+rand.Intn(1000))).Err()
	if err != nil {
		return err
	}
	return nil
}

func (g *goodsRepo) DeleteGoodsCache(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	err := g.data.rc.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
