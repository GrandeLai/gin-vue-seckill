package main

import (
	"fmt"
	"gin-vue-seckill/jdaw_web/common/utils"
	"github.com/go-redis/redis"
	"time"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

func main() {
	client = redis.NewClient(&redis.Options{
		Addr:     "192.168.10.3:6379", // no password set
		DB:       0,                   // use default DB
		PoolSize: 100,
	})

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	map_data := map[string]interface{}{
		"code": 500,
		"msg":  "下单失败",
	}
	mapd := utils.MapToStr(map_data)
	client.Set("sjsj", mapd, 60*time.Second)
	f, err := client.Get("sjsj").Result()
	fmt.Println(f)
}
