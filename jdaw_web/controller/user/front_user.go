package user

import (
	"context"
	"gin-vue-seckill/jdaw_user_srv/proto/front_user"
	"gin-vue-seckill/jdaw_web/common/utils"
	"github.com/asim/go-micro/plugins/client/grpc/v3"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Sendmail(ctx *gin.Context) {
	email := ctx.PostForm("email")
	if !utils.VerifyEmailFormat(email) {
		ctx.JSON(http.StatusHTTPVersionNotSupported, gin.H{
			"code": 505,
			"msg":  "邮箱格式不正确",
		})
	} else {
		//配置注册中心
		consulReg := consul.NewRegistry(
			registry.Addrs("127.0.0.1:8500"),
		)
		s := micro.NewService(
			micro.Registry(consulReg), //设置注册中心
			micro.Client(grpc.NewClient()),
		)
		//)
		s.Init()
		client := front_user.NewFrontUserService("jdaw.user.srv", s.Client())
		rep, _ := client.FrontUserSendEmail(context.TODO(), &front_user.FrontUserMailRequest{Mail: email})

		ctx.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
	}
}

func FrontUserRegister(ctx *gin.Context) {
	email := ctx.PostForm("email")
	if utils.VerifyEmailFormat(email) {
		code := ctx.PostForm("captche")
		password := ctx.PostForm("password")
		repassword := ctx.PostForm("repassword")
		if password != repassword {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "两次输入的密码不一致，请重新输入",
			})
		} else {
			//配置注册中心
			consulReg := consul.NewRegistry(
				registry.Addrs("127.0.0.1:8500"),
			)

			s := micro.NewService(
				micro.Registry(consulReg), //设置注册中心
				micro.Client(grpc.NewClient()),
			)
			s.Init()
			client := front_user.NewFrontUserService("jdaw.user.srv", s.Client())
			rep, err := client.FrontUserRegister(context.TODO(), &front_user.FrontUserRequest{
				Email:      email,
				Code:       code,
				Password:   password,
				Repassword: repassword,
			})
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code": rep.Code,
					"msg":  rep.Msg,
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code": rep.Code,
					"msg":  rep.Msg,
				})
			}
		}
	} else {
		ctx.JSON(http.StatusHTTPVersionNotSupported, gin.H{
			"code": 505,
			"msg":  "邮箱格式不正确",
		})
	}
}

func FrontUserLogin(ctx *gin.Context) {
	mail := ctx.PostForm("mail")
	password := ctx.PostForm("password")

	if is_ok := utils.VerifyEmailFormat(mail); is_ok {
		//配置注册中心
		consulReg := consul.NewRegistry(
			registry.Addrs("127.0.0.1:8500"),
		)

		s := micro.NewService(
			micro.Registry(consulReg), //设置注册中心
			micro.Client(grpc.NewClient()),
		)
		s.Init()
		client := front_user.NewFrontUserService("jdaw.user.srv", s.Client())
		rep, err := client.FrontUserLogin(ctx, &front_user.FrontUserRequest{
			Email:    mail,
			Password: password,
		})
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": rep.Code,
				"msg":  rep.Msg,
			})
		} else {
			token, err1 := utils.GenToken(rep.UserName, utils.FrontUserExpireDuration, utils.FrontUserSecretKey)
			if err1 != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code": rep.Code,
					"msg":  rep.Msg,
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code":     rep.Code,
					"msg":      rep.Msg,
					"token":    token,
					"username": mail,
				})
			}
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "邮箱格式不正确",
		})
	}

}
