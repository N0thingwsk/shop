package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"shop/app/user/internal/biz"
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
func (r *userRepo) GetUserinfo(ctx context.Context, user *biz.User) (*biz.User, error) {

	result := r.data.db.Where("id = ?", user.ID).First(&user)
	if result.RowsAffected == 0 {
		return &biz.User{}, errors.New("GetUserinfo err: user is nil")
	}
	if result.Error != nil {
		return &biz.User{}, errors.New("GetUserinfo err: user get fail")
	}
	return user, nil
}

func (r *userRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	result := r.data.db.Where("mobile = ?", user.Mobile)
	if result.RowsAffected == 1 {
		log.Error("CreateUser err: user already exists")
		return &biz.User{}, errors.New("user already exists")
	}
	t := time.Now()
	result = r.data.db.Create(&biz.User{
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: fmt.Sprintf("用户%s", user.Mobile),
		Birthday: &t,
		Gender:   "男",
	})
	if result.Error != nil {
		return &biz.User{}, result.Error
	}
	return &biz.User{
		Mobile:   user.Mobile,
		NickName: fmt.Sprintf("用户%s", user.Mobile),
		Birthday: &t,
		Gender:   "男",
	}, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	result := r.data.db.Model(&biz.User{}).Where("id = ?", user.ID).Update("nick_name", user.NickName)
	if result.Error != nil {
		return &biz.User{}, errors.New("NickName更新失败")
	}
	return user, nil
}

func (r *userRepo) DeleteUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	result := r.data.db.Delete(&biz.User{}, id)
	if result.Error != nil {

	}
	return &biz.User{}, nil
}
