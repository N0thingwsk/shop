package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"shop/app/user/internal/biz"
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

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (r *userRepo) GetUserList(ctx context.Context, page int) ([]biz.User, error) {
	var users []biz.User
	result := r.data.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	//rspTotal := result.RowsAffected
	r.data.db.Scopes(Paginate(page, 1)).Find(&users)
	return nil, nil
}

func (r *userRepo) GetUserinfoByMobile(ctx context.Context, mobile string) (biz.User, error) {
	var user biz.User
	result := r.data.db.Where(&biz.User{Mobile: mobile}).First(&user)
	if result.RowsAffected == 0 {
		return biz.User{}, errors.New("用户不存在")
	}
	if result.Error != nil {
		return biz.User{}, result.Error
	}
	return user, nil
}

func (r *userRepo) GetUserinfoById(context.Context, int) (biz.User, error) {
	return biz.User{}, nil
}

func (r *userRepo) CreateUser(context.Context, biz.User) (biz.User, error) {
	return biz.User{}, nil
}

func (r *userRepo) UpdateUser(context.Context, int) (biz.User, error) {
	return biz.User{}, nil
}

func (r *userRepo) DeleteUser(context.Context, int) (biz.User, error) {
	return biz.User{}, nil
}