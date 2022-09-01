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

func (r *userRepo) GetUserinfo(ctx context.Context, id int) (biz.User, error) {
	var user biz.User
	result := r.data.db.Where("id = ?", id).First(&user)
	if result.RowsAffected == 0 {
		return biz.User{}, errors.New("用户不存在")
	}
	if result.Error != nil {
		return biz.User{}, errors.New("用户信息获取失败")
	}
	return user, nil
}

func (r *userRepo) LoginUser(ctx context.Context, use biz.User) (biz.User, error) {
	var user biz.User
	result := r.data.db.Where("mobile = ?", use.Mobile).First(&user)
	if result.RowsAffected == 0 {
		return biz.User{}, errors.New("用户不存在")
	}
	if result.Error != nil {
		return biz.User{}, errors.New("登录失败")
	}
	return user, nil
}

func (r *userRepo) CreateUser(ctx context.Context, use biz.User) (biz.User, error) {
	result := r.data.db.Where("mobile = ?", use.Mobile).First(&use)
	if result.RowsAffected == 1 {
		return biz.User{}, errors.New("用户已存在")
	}
	t := time.Now()
	result = r.data.db.Create(&biz.User{
		Mobile:   use.Mobile,
		Password: use.Password,
		NickName: fmt.Sprintf("用户%s", use.Mobile),
		Birthday: &t,
		Gender:   "男",
	})
	if result.Error != nil {
		return biz.User{}, result.Error
	}
	return biz.User{
		Mobile:   use.Mobile,
		NickName: fmt.Sprintf("用户%s", use.Mobile),
		Birthday: &t,
		Gender:   "男",
	}, nil
}

func (r *userRepo) UpdateUserNickName(ctx context.Context, use biz.User) (biz.User, error) {
	result := r.data.db.Model(&biz.User{}).Where("id = ?", use.ID).Update("nick_name", use.NickName)
	if result.Error != nil {
		return biz.User{}, errors.New("NickName更新失败")
	}
	return use, nil
}

func (r *userRepo) UpdateUserPassword(ctx context.Context, use biz.User) (biz.User, error) {
	result := r.data.db.Model(&biz.User{}).Where("id = ?", use.ID).Update("password", use.Password)
	if result.Error != nil {
		return biz.User{}, errors.New("PassWord更新失败")
	}
	return use, nil
}

func (r *userRepo) UpdateUserBirthday(ctx context.Context, use biz.User) (biz.User, error) {
	result := r.data.db.Model(&biz.User{}).Where("id = ?", use.ID).Update("birthday", use.Birthday)
	if result.Error != nil {
		return biz.User{}, errors.New("BirthDay更新失败")
	}
	return use, nil
}

func (r *userRepo) UpdateUserGender(ctx context.Context, use biz.User) (biz.User, error) {
	result := r.data.db.Model(&biz.User{}).Where("id = ?", use.ID).Update("gender", use.Gender)
	if result.Error != nil {
		return biz.User{}, errors.New("Gender更新失败")
	}
	return use, nil
}

func (r *userRepo) DeleteUser(ctx context.Context, id int) (biz.User, error) {
	r.data.db.Delete(&biz.User{}, id)
	return biz.User{}, nil
}
