package user

import (
	"gin-vue-seckill/jdaw_web/common/middleware"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {

	router.POST("/send_email", Sendmail)
	//注册
	router.POST("/front_user_register", FrontUserRegister)

	//前端用户登录
	router.POST("/front_user_login", FrontUserLogin)

	//管理员登录
	router.POST("/admin_login", AdminUserLogin)
	router.GET("/get_front_users", middleware.JstTokenValid, GetFrontUserList)
}
