package seckill

import (
	"fmt"
	"gin-vue-seckill/jdaw_web/common/utils"
	"gin-vue-seckill/jdaw_web/rabbitmq"
	redis "gin-vue-seckill/jdaw_web/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SecKill(ctx *gin.Context) {
	id := ctx.PostForm("id")
	front_user_email, exist := ctx.Get("front_user_name")
	if !exist {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "token有问题",
		})
	}
	//consulReg := consul.NewRegistry(
	//	registry.Addrs("127.0.0.1:8500"),
	//)
	//
	//s := micro.NewService(
	//	micro.Registry(consulReg), //设置注册中心
	//	micro.Client(grpc.NewClient()),
	//)
	//s.Init()
	//client := seckills.NewSecKillService("jdaw.seckill.srv", s.Client())
	//rep, _ := client.FrontSecKill(ctx, &seckills.SecKillRequest{
	//	Id:             int32(utils.StrToInt(id)),
	//	FrontUserEmail: front_user_email.(string),
	//})

	qe := rabbitmq.QueueAndExchange{
		QueueName:    "jdaw_web.order_queue",
		ExchangeName: "jdaw_web.order_exchange",
		ExchangeType: "direct",
		RoutingKey:   "jdaw_web.order",
	}
	mq := rabbitmq.NewRabbitMq(qe)
	mq.ConnMq()
	mq.OpenChannel()
	defer func() {
		mq.CloseMq()
	}()
	defer func() {
		mq.CloseChannel()
	}()
	order_map := map[string]interface{}{
		"uemail": front_user_email,
		"pid":    id,
	}
	mq.PublishMsg(utils.MapToStr(order_map))

	ctx.JSON(http.StatusOK, gin.H{
		"code": 500,
		"msg":  "下单中，请稍后",
	})
}

func GetSeckillResult(ctx *gin.Context) {

	uemail, exist := ctx.Get("front_user_name")

	if exist {
		ret, err_r := redis.Client.Get(uemail.(string)).Result()

		if err_r == nil {
			ret_map := utils.StrToMap(ret)
			fmt.Println(ret_map)
			ctx.JSON(http.StatusOK, gin.H{ // 说明从redis里面获取到了数据，
				"code": 200,
				"msg":  ret_map["msg"],
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
		})
		return
	}

}
