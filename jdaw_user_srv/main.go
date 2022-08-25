package main

import (
	"fmt"
	"gin-vue-seckill/jdaw_user_srv/common/utils/snowflake"
	"gin-vue-seckill/jdaw_user_srv/controller"
	_ "gin-vue-seckill/jdaw_user_srv/data_source"
	"gin-vue-seckill/jdaw_user_srv/proto/admin_user"
	"gin-vue-seckill/jdaw_user_srv/proto/front_user"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/plugins/server/grpc/v3"
	service "github.com/asim/go-micro/v3" // "github.com/micro/micro/v3/service"
	"github.com/asim/go-micro/v3/logger"  // "github.com/micro/micro/v3/service/logger"
	"github.com/asim/go-micro/v3/registry"
)

const (
	ServerName = "jdaw.user.srv"
	ConsulAddr = "127.0.0.1:8500"
)

func main() {
	// 新建注册
	consulReg := consul.NewRegistry(
		registry.Addrs(ConsulAddr),
	)

	grpcServer := grpc.NewServer()
	srv := service.NewService(
		service.Server(grpcServer),
		service.Name(ServerName),    // 服务名字
		service.Registry(consulReg), // 注册中心
	)

	//注册handler
	_ = front_user.RegisterFrontUserHandler(srv.Server(), new(controller.Front_user))
	_ = admin_user.RegisterAdminUserHandler(srv.Server(), new(controller.Admin_user))

	//// 加载配置
	//if err := settings.Init(); err != nil {
	//	fmt.Printf("init settings failed, err:%v\n", err)
	//	return
	//}
	//fmt.Println(settings.Conf.StartTime)
	//// 初始化日志
	//if err := log.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
	//	fmt.Printf("init logger failed, err:%v\n", err)
	//	return
	//}
	//defer zap.L().Sync()
	//zap.L().Debug("logger init success...")

	//雪花算法初始化
	if err := snowflake.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
