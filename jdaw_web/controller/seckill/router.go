package seckill

import (
	"gin-vue-seckill/jdaw_web/common/middleware"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	// 秒杀接口
	router.POST("/front/seckill", middleware.JwtTokenFrontValid, SecKill)
	router.GET("/front/get_seckill_result", middleware.JwtTokenFrontValid, GetSeckillResult)
	//前端界面
	router.GET("/front/get_seckill_list", GetFrontSeckillList)
	router.GET("/front/seckill_detail", middleware.JwtTokenFrontValid, GetSeckillInfo)
	router.GET("/get_seckill_list", middleware.JstTokenValid, GetSeckillList)
	router.GET("/get_products", middleware.JstTokenValid, GetProduct)
	router.POST("/seckill_del", middleware.JstTokenValid, DeleteProduct)
	router.POST("/seckill_add", middleware.JstTokenValid, AddSeckill)
	router.GET("/seckill_to_edit", middleware.JstTokenValid, ToEditSeckill)
	router.POST("/seckill_do_edit", middleware.JstTokenValid, UpdateSeckill)

}
