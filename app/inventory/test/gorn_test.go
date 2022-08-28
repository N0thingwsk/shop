package test

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	v1 "shop/api/inventory/v1"
	"shop/app/inventory/internal/biz"
	"sync"
	"testing"
)

func TestGrom(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/shop_inventory?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i <= 10; i++ {
		db.Create(&biz.Inventory{
			Goods:  int32(i),
			Stocks: int32(i),
		})
	}
}

func TestConsul(t *testing.T) {
	consulCli, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	r := consul.New(consulCli)
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///inventory"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	c := v1.NewInventoryClient(conn)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("启动协程")
			detail, err := c.Sell(context.Background(), &v1.SellInfo{
				GoodsInfo: []*v1.GoodsInvInfo{
					{Id: 1, Stocks: 10},
					{Id: 2, Stocks: 10},
				},
			})
			if err != nil {
				log.Error(err.Error())
				wg.Done()
				return
			}
			fmt.Println(detail)
			wg.Done()
			fmt.Println("执行成功")
		}()
	}
	wg.Wait()
	//config := api.DefaultConfig()
	//config.Address = "127.0.0.1:8500"
	//client, err := api.NewClient(config)
	//if err != nil {
	//	fmt.Println("consul client error : ", err)
	//	return
	//}
	//services, err := client.Agent().Services()
	//if err != nil {
	//	return
	//}
	//fmt.Println(services["const.local"].Address, services["const.local"].Port)
	//conn, err := grpc.Dial("192.168.3.146:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer conn.Close()
	//c := v1.NewInventoryClient(conn)
	//detail, err := c.InvDetail(context.Background(), &v1.GoodsInvInfo{
	//	Id: 1,
	//})
	//if err != nil {
	//	return
	//}
	//fmt.Println(detail)
}

func TestGo(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/shop_inventory?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err.Error())
	}
	go func() {
		result := db.Exec("update inventory set stocks = stocks + 2, version = version = +1 where goods = 1 and version = version;")
		if result.Error != nil {
			panic(err.Error())
		}
	}()
	go func() {
		result := db.Exec("update inventory set stocks = stocks + 2, version = version = +1 where goods = 1 and version = version;")
		if result.Error != nil {
			panic(err.Error())
		}
	}()
}

func TestSsa(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	for _, x := range arr {
		fmt.Println(x)
		if x == 0 {
			break
		} else {
			fmt.Println(1)
			continue
		}
	}
}
