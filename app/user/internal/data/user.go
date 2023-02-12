package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v9"
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
	cache, err := r.data.rc.Get(ctx, strconv.Itoa(uid)).Bytes()
	if err == redis.Nil {
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
	} else if err != nil {
		return biz.User{}, err
	}
	if err := r.data.rc.Expire(ctx, "1", time.Second*60).Err(); err != nil {
		log.Error("GetUserinfo err: redis time reset fail")
	}
	var info v1.UserInfoReply
	err = proto.Unmarshal(cache, &info)
	if err != nil {
		return biz.User{}, errors.New("GetUserinfo err: protobuf unmarshal err")
	}
	return biz.User{
		Mobile:   info.User.GetMobile(),
		NickName: info.User.GetNikeName(),
		Birthday: info.User.GetBirthday(),
		Gender:   info.User.GetGender(),
	}, nil
}

func (r *userRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	tx := r.data.db.Begin()
	result := tx.Where("mobile = ?", user.Mobile)
	if result.RowsAffected == 1 {
		log.Error("CreateUser err: user already exists")
		tx.Rollback()
		return &biz.User{}, errors.New("user already exists")
	}
	result = tx.Create(&biz.User{
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: fmt.Sprintf("用户%s", user.Mobile),
		Birthday: time.Now().String(),
		Gender:   "男",
	})
	if result.Error != nil {
		tx.Rollback()
		return &biz.User{}, result.Error
	}
	tx.Commit()
	return biz.User{
		Mobile:   user.Mobile,
		NickName: fmt.Sprintf("用户%s", user.Mobile),
		Birthday: user.Birthday,
		Gender:   "男",
	}, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	tx := r.data.db.Begin()
	result := tx.Model(&biz.User{}).Where("id = ?", user.ID).Updates(&user)
	if result.Error != nil {
		tx.Rollback()
		return biz.User{}, errors.New("NickName更新失败")
	}
	tx.Commit()
	return user, nil
}

func (r *userRepo) DeleteUser(ctx context.Context, uid int) (biz.User, error) {
	result := r.data.db.Delete(&biz.User{}, uid)
	if result.Error != nil {

	}
	return biz.User{}, nil
}
