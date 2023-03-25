package data

import (
	"context"
	"time"
)

// GetUserCache 获取用户缓存
func (r *userRepo) GetUserCache(ctx context.Context, uid string) ([]byte, error) {
	cache, err := r.data.rc.Get(ctx, uid).Bytes()
	if err != nil {
		r.log.Error("GetUserCache err: redis get err", err.Error())
		return nil, err
	}
	err = r.data.rc.Expire(ctx, uid, 60*time.Second).Err()
	if err != nil {
		r.log.Error("GetUserCache err: redis expire err", err.Error())
	}
	return cache, nil
}

// SetUserCache 设置用户缓存
func (r *userRepo) SetUserCache(ctx context.Context, uid string, user string) error {
	err := r.data.rc.Set(ctx, uid, user, 60*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

// DelUserCache 删除用户缓存
func (r *userRepo) DelUserCache(ctx context.Context, uid string) error {
	err := r.data.rc.Del(ctx, uid).Err()
	if err != nil {
		return err
	}
	return nil
}
