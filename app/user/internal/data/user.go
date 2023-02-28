package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/proto"
	v1 "shop/api/user/v1"
	"shop/app/user/internal/biz"
	"strconv"
	"time"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// GetUserinfo 获取用户信息
func (r *userRepo) GetUserinfo(ctx context.Context, uid int) (biz.User, error) {
	user := biz.User{}
	result := r.data.db.Where("id = ?", uid).First(&user)
	if result.RowsAffected == 0 {
		return biz.User{}, errors.New("GetUserinfo err: user is nil")
	}
	if result.Error != nil {
		return biz.User{}, errors.New("GetUserinfo err: user get fail")
	}
	info := v1.UserInfoReply{
		User: &v1.UserInfoReply_User{
			Mobile:   user.Mobile,
			NikeName: user.NickName,
			Birthday: user.Birthday,
			Gender:   user.Gender,
		},
	}
	marshal, err := proto.Marshal(&info)
	if err != nil {
		log.Error("GetUserinfo err: protobuf marshal err")
	}
	err = r.data.rc.Set(ctx, strconv.Itoa(int(user.ID)), marshal, time.Second*60).Err()
	if err != nil {
		log.Error("GetUserinfo err: redis write err", err.Error())
	}
	return user, nil
}

// GetUserCache 获取用户缓存
func (r *userRepo) GetUserCache(ctx context.Context, key string) ([]byte, error) {
	cache, err := r.data.rc.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	if err := r.data.rc.Expire(ctx, key, time.Second*60).Err(); err != nil {
		log.Error("GetUserinfo err: redis time reset fail")
	}
	return cache, nil
}

func (r *userRepo) GetUserInfoUserName(c) {

}

// CreateUser 创建用户
func (r *userRepo) CreateUser(ctx context.Context, user biz.User) error {
	tx := r.data.db.Begin()
	result := tx.Create(&biz.User{
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: fmt.Sprintf("用户%s", user.Mobile),
		Birthday: time.Now().String(),
		Gender:   "男",
	})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

// UpdateUser 更新用户
func (r *userRepo) UpdateUser(ctx context.Context, user biz.User) error {
	tx := r.data.db.Begin()
	result := tx.Model(&biz.User{}).Where("id = ?", user.ID).Updates(&user)
	if result.Error != nil {
		tx.Rollback()
		return errors.New("update err:" + result.Error.Error())
	}
	tx.Commit()
	return nil
}

// DeleteUser 删除用户
func (r *userRepo) DeleteUser(ctx context.Context, uid int) error {
	result := r.data.db.Delete(&biz.User{}, uid)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
