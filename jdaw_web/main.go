package main

import (
	"fmt"
	"gin-vue-seckill/jdaw_web/common/middleware"
	redis "gin-vue-seckill/jdaw_web/redis"
	"gin-vue-seckill/jdaw_web/router"
	"github.com/gin-gonic/gin"
)

const addr = ":8081"

func main() {
	r := gin.Default()
	//redis初始化
	if err_redis := redis.Init(); err_redis != nil {
		fmt.Println(err_redis)
	}
	//使用全局中间件处理跨域请求
	r.Use(middleware.CrosMiddleWare)
	router.InitRouter(r)
	if err := r.Run(addr); err != nil {
		fmt.Println("err")
	}
}
