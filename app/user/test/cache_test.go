package test

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v9"
	"google.golang.org/protobuf/proto"
	v1 "shop/api/user/v1"
	"testing"
)

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
	set := rc.Set(context.Background(), "1", string(marshal), 0).Err()
	fmt.Println(set)
	result, err := rc.Get(context.Background(), "1").Result()
	if err != nil {
		return
	}
	fmt.Println(result)
}
