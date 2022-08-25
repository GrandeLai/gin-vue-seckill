package user

import (
	"gin-vue-seckill/jdaw_user_srv/proto/admin_user"
	"gin-vue-seckill/jdaw_web/common/utils"
	"github.com/asim/go-micro/plugins/client/grpc/v3"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AdminUserLogin(ctx *gin.Context) {
	userName := ctx.PostForm("username")
	password := ctx.PostForm("password")

	//配置注册中心
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	s := micro.NewService(
		micro.Registry(consulReg), //设置注册中心
		micro.Client(grpc.NewClient()),
	)
	s.Init()
	client := admin_user.NewAdminUserService("jdaw.user.srv", s.Client())
	rep, err := client.AdminUserLogin(ctx, &admin_user.AdminUserRequest{
		Username: userName,
		Password: password,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
	} else {
		admin_token, err1 := utils.GenToken(rep.UserName, utils.AdminUserExpireDuration, utils.AdminUserSecretKey)
		if err1 != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": rep.Code,
				"msg":  rep.Msg,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":        rep.Code,
				"msg":         rep.Msg,
				"admin_token": admin_token,
				"username":    rep.UserName,
			})
		}
	}
}

func GetFrontUserList(ctx *gin.Context) {
	cp := ctx.Query("currentPage")
	ps := ctx.Query("pageSize")
	currentPage, _ := strconv.ParseInt(cp, 10, 0)
	pageSize, _ := strconv.ParseInt(ps, 10, 0)

	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	s := micro.NewService(
		micro.Registry(consulReg), //设置注册中心
		micro.Client(grpc.NewClient()),
	)
	s.Init()
	client := admin_user.NewAdminUserService("jdaw.user.srv", s.Client())
	rep, err := client.GetFrontUserList(ctx, &admin_user.FrontUserListRequest{
		CurrentPage: int32(currentPage),
		PageSize:    int32(pageSize),
	})

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":         rep.Code,
			"msg":          rep.Msg,
			"front_users":  rep.Frontuserlist,
			"total":        int(rep.Total),
			"current_page": int(rep.Current),
			"page_size":    int(rep.PageSize),
		})
	}

}
