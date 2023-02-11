package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"shop/app/user/internal/conf"
)

//var (
//	// ErrUserNotFound is user not found.
//	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
//)

type User struct {
	gorm.Model
	Mobile   string `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	NickName string `gorm:"type:varchar(20)"`
	Birthday string `gorm:"type:varchar(20)"`
	Gender   string `gorm:"column:gender;default:male;type:varchar(6)"`
	Role     int    `gorm:"column:role;default:1;type:int"`
}

type UserRepo interface {
	GetUserinfo(context.Context, User) (User, error)
	CreateUser(context.Context, User) (User, error)
	UpdateUser(context.Context, User) (User, error)
	DeleteUser(context.Context, User) (User, error)
}

type UserUsecase struct {
	repo UserRepo
	c    *conf.Server
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, c *conf.Server, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, c: c, log: log.NewHelper(logger)}
}

func (u *UserUsecase) GetUserinfo(ctx context.Context, id int) (User, error) {
	user, err := u.repo.GetUserinfo(ctx, User{Model: gorm.Model{ID: uint(id)}})
	if err != nil {
		return User{}, err
	}
	return user, nil
}

//func (u *UserUsecase) LoginUser(ctx context.Context, user User) (string, error) {
//	use, err := u.repo.LoginUser(ctx, user)
//	if err != nil {
//		return "", err
//	}
//	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
//	passwordInfo := strings.Split(use.Password, "$")
//	isPass := password.Verify(user.Password, passwordInfo[2], passwordInfo[3], options)
//	if isPass != true {
//		return "", errors.New("密码错误")
//	}
//	return auth.GenerateToken(u.c.Token.Secret, use.ID), nil
//}
//
//func (u *UserUsecase) CreateUser(ctx context.Context, user User) (User, error) {
//	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
//	salt, encodedPwd := password.Encode(user.Password, options)
//	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
//	us, err := u.repo.CreateUser(ctx, User{
//		Mobile:   user.Mobile,
//		Password: newPassword,
//		NickName: user.NickName,
//	})
//	if err != nil {
//		return User{}, err
//	}
//	return User{
//		Mobile:   us.Mobile,
//		NickName: us.NickName,
//	}, nil
//}
//
//func (u *UserUsecase) UpdateUserNickName(ctx context.Context, user User) (User, error) {
//	result, err := u.repo.UpdateUserNickName(ctx, user)
//	if err != nil {
//		return User{}, err
//	}
//	return result, nil
//}
//
//func (u *UserUsecase) UpdateUserPassword(ctx context.Context, user User) (User, error) {
//	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
//	salt, encodedPwd := password.Encode(user.Password, options)
//	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
//	user.Password = newPassword
//	result, err := u.repo.UpdateUserPassword(ctx, user)
//	if err != nil {
//		return User{}, err
//	}
//	return result, nil
//}
//
//func (u *UserUsecase) UpdateUserBirthday(ctx context.Context, user User) (User, error) {
//	result, err := u.repo.UpdateUserBirthday(ctx, user)
//	if err != nil {
//		return User{}, err
//	}
//	return result, nil
//}
//
//func (u *UserUsecase) UpdateUserGender(ctx context.Context, user User) (User, error) {
//	result, err := u.repo.UpdateUserGender(ctx, user)
//	if err != nil {
//		return User{}, err
//	}
//	return result, nil
//}
//
//func (u *UserUsecase) DeleteUser(ctx context.Context, id int) (User, error) {
//	us, err := u.DeleteUser(ctx, id)
//	if err != nil {
//		return User{}, err
//	}
//	return us, nil
//}
