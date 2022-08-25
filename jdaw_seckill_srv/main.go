package main

import (
	"gin-vue-seckill/jdaw_seckill_srv/controller"
	_ "gin-vue-seckill/jdaw_seckill_srv/data_source"
	seckills "gin-vue-seckill/jdaw_seckill_srv/proto/seckill"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/plugins/server/grpc/v3"
	service "github.com/asim/go-micro/v3" // "github.com/micro/micro/v3/service"
	"github.com/asim/go-micro/v3/logger"  // "github.com/micro/micro/v3/service/logger"
	"github.com/asim/go-micro/v3/registry"
)

const (
	ServerName = "jdaw.seckill.srv"
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

	_ = seckills.RegisterSecKillHandler(srv.Server(), new(controller.Seckill))

	////redis初始化
	//if err_redis := data_source.Init(); err_redis != nil {
	//	fmt.Println(err_redis)
	//}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
