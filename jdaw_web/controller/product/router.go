package product

import (
	"gin-vue-seckill/jdaw_web/common/middleware"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	router.Use(middleware.JstTokenValid)
	router.GET("/get_product_list", GetProductList)
	router.POST("/product_add", AddProduct)
	router.POST("/product_del", DeleteProduct)
	router.GET("/to_product_edit", ToProductEdit)
	router.POST("/do_product_edit", UpdataProduct)
}
