package router

import (
	"gin-vue-seckill/jdaw_web/controller/product"
	"gin-vue-seckill/jdaw_web/controller/seckill"
	"gin-vue-seckill/jdaw_web/controller/user"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	user_group := router.Group("/user")
	product_group := router.Group("/product")
	seckill_group := router.Group("/seckill")

	user.Router(user_group)
	product.Router(product_group)
	seckill.Router(seckill_group)
}
