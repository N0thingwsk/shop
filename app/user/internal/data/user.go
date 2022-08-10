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

func (r *userRepo) GetUserinfoById(ctx context.Context, id int) (biz.User, error) {
	var user biz.User
	result := r.data.db.First(&user, id)
	if result.RowsAffected == 0 {
		return biz.User{}, errors.New("用户不存在")
	}
	if result.Error != nil {
		return biz.User{}, result.Error
	}
	return user, nil
}

func (r *userRepo) CreateUser(ctx context.Context, use biz.User) (biz.User, error) {
	result := r.data.db.Where("mobile= ?", use.Mobile).First(&use)
	if result.RowsAffected == 1 {
		return biz.User{}, errors.New("用户已存在")
	}
	result = r.data.db.Create(&biz.User{
		Mobile:   use.Mobile,
		Password: use.Password,
		NickName: use.NickName,
	})
	if result.Error != nil {
		return biz.User{}, result.Error
	}
	return biz.User{
		Mobile:   use.Mobile,
		Password: use.Password,
		NickName: use.NickName,
	}, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, use biz.User) (biz.User, error) {
	result := r.data.db.Save(&use)
	if result.Error != nil {
		return biz.User{}, result.Error
	}
	return use, nil
}

func (r *userRepo) DeleteUser(ctx context.Context, id int) (biz.User, error) {
	r.data.db.Delete(&biz.User{}, id)
	return biz.User{}, nil
}
