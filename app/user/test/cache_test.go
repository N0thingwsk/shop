package test

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v9"
	"google.golang.org/protobuf/proto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	v1 "shop/api/user/v1"
	"shop/app/user/internal/biz"
	"strconv"
	"testing"
	"time"
)

func wirteredis(marshal []byte, rc *redis.Client) {
	err := rc.Set(context.Background(), "1", string(marshal), time.Second*60).Err()
	if err != nil {
		log.Error(err.Error())
	}
}

func TestCache(t *testing.T) {
	info := v1.UserInfoReply{
		User: &v1.UserInfoReply_User{
			Mobile:   "12345678910",
			NikeName: "Test",
			Birthday: "2000/11/23",
			Gender:   "男",
		},
	}
	marshal, err := proto.Marshal(&info)
	if err != nil {
		log.Error("GetUserinfo err: protobuf marshal err")
	}
	rc := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	if result, err := rc.Ping(context.Background()).Result(); err != nil {
		log.Debug(result)
	}
	wirteredis(marshal, rc)
	result, err := rc.Get(context.Background(), "1").Bytes()
	if err != nil {
		panic(err.Error())
	} else if err == redis.Nil {
		log.Error(err.Error())
	}
	var v v1.UserInfoReply
	err = proto.Unmarshal(result, &v)
	if err != nil {
		log.Error(err.Error())
	}
	fmt.Println(v.User.Mobile)
}

func TestMysql(t *testing.T) {
	addr := "root:123456@tcp(127.0.0.1:3307)/shop_user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&biz.User{})
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i <= 20; i++ {
		db.Create(&biz.User{
			Mobile:   "111111111" + strconv.Itoa(i),
			Password: "test" + strconv.Itoa(i),
			NickName: "test" + strconv.Itoa(i),
			Birthday: time.Now().String(),
			Gender:   "男",
			Role:     i,
		})
	}
}
