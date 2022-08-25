package controller

import (
	"context"
	"fmt"
	"gin-vue-seckill/jdaw_seckill_srv/common/utils"
	"gin-vue-seckill/jdaw_seckill_srv/data_source"
	"gin-vue-seckill/jdaw_seckill_srv/models"
	"gin-vue-seckill/jdaw_seckill_srv/proto/seckill"
	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
	"time"
)

type Seckill struct{}

func (s *Seckill) FrontSecKill(ctx context.Context, in *seckill.SecKillRequest, out *seckill.SecKillResponse) error {
	id := in.Id
	front_user_email := in.FrontUserEmail
	seckill := models.SecKills{}
	now_time := time.Now()
	result := data_source.Db.Where("id=?", id).Find(&seckill)
	if result.Error != nil {
		out.Code = 500
		out.Msg = "下单失败"
		return nil
	}
	ret_num := result.Where("num>0").Find(&seckill)
	if ret_num.Error != nil {
		out.Code = 500
		out.Msg = "商品已被抢完"
		return nil
	}
	re_time := ret_num.Where("start_time <= ?", now_time).Where("end_time > ?", now_time).Find(&seckill)
	if re_time.Error != nil {
		out.Code = 500
		out.Msg = "不在抢购时间内"
		return nil
	}

	order_re := models.Orders{}
	ret_eamil := data_source.Db.Where("uemail=?", front_user_email).Where("sid=?", id).Find(&order_re)
	if ret_eamil.Error == nil {
		out.Code = 500
		out.Msg = "每个用户只能下单一次"
		return nil
	}

	ret := re_time.Update("num", seckill.Num-1)
	if ret.Error != nil {
		out.Code = 500
		out.Msg = "下单失败"
		return nil
	}
	//用户只能购买一个
	order := models.Orders{
		Uemail: front_user_email,
		Sid:    int(id),
	}
	ret_order := data_source.Db.Create(&order)
	if ret_order.Error != nil {
		out.Code = 500
		out.Msg = "下单失败"
		return nil
	}
	out.Code = 200
	out.Msg = "下单成功"
	return nil
}

//从队列中获取任务消费
func init() {
	conn, err := amqp.Dial("amqp://admin:admin@192.168.10.3:5672")
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	ch.Qos(10, 0, false)
	deleveries, err := ch.Consume("jdaw_web.order_queue", "order_consumer", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
	}
	for delivery := range deleveries {
		go OrderApply(delivery)
	}
}

func OrderApply(delivery amqp.Delivery) {
	Client := redis.NewClient(&redis.Options{
		Addr:     "192.168.10.3:6379", // no password set
		DB:       0,                   // use default DB
		PoolSize: 100,
	})
	body := delivery.Body
	request_data := utils.StrToMap(string(body))

	id := request_data["pid"]
	front_user_email := request_data["uemail"].(string)

	seckill := models.SecKills{}
	now_time := time.Now()
	result := data_source.Db.Where("id=?", id).Find(&seckill)

	if result.Error != nil {
		delivery.Ack(true)
		map_data := map[string]interface{}{
			"code": 500,
			"msg":  "下单失败",
		}
		Client.Set(front_user_email, utils.MapToStr(map_data), 1*time.Hour)
		return
	}
	ret_num := result.Where("num>0").Find(&seckill)
	if ret_num.Error != nil {
		delivery.Ack(true)
		map_data := map[string]interface{}{
			"code": 500,
			"msg":  "商品已被抢完",
		}
		Client.Set(front_user_email, utils.MapToStr(map_data), 1*time.Hour)
		return
	}
	re_time := ret_num.Where("start_time <= ?", now_time).Where("end_time > ?", now_time).Find(&seckill)
	if re_time.Error != nil {
		delivery.Ack(true)
		map_data := map[string]interface{}{
			"code": 500,
			"msg":  "不在抢购时间内",
		}
		data := utils.MapToStr(map_data)
		Client.Set(front_user_email, data, 1*time.Hour)
		return
	}

	order_re := models.Orders{}
	ret_eamil := data_source.Db.Where("uemail=?", front_user_email).Where("sid=?", id).Find(&order_re)
	if ret_eamil.Error == nil {
		delivery.Ack(true)
		map_data := map[string]interface{}{
			"code": 500,
			"msg":  "不要重复下单",
		}
		Client.Set(front_user_email, utils.MapToStr(map_data), 2*time.Hour)
		return
	}

	ret := re_time.Update("num", seckill.Num-1)
	if ret.Error != nil {
		delivery.Ack(true)
		map_data := map[string]interface{}{
			"code": 500,
			"msg":  "下单失败",
		}
		Client.Set(front_user_email, utils.MapToStr(map_data), 2*time.Hour)
		return
	}
	//用户只能购买一个
	order := models.Orders{
		Uemail:     front_user_email,
		Sid:        utils.StrToInt(id.(string)),
		CreateTime: time.Now(),
	}
	ret_order := data_source.Db.Create(&order)
	if ret_order.Error != nil {
		delivery.Ack(true)
		map_data := map[string]interface{}{
			"code": 500,
			"msg":  "下单失败",
		}
		Client.Set(front_user_email, utils.MapToStr(map_data), 1*time.Hour)
		return
	}
	delivery.Ack(true)
	map_data := map[string]interface{}{
		"code": 200,
		"msg":  "下单成功",
	}
	Client.Set(front_user_email, utils.MapToStr(map_data), 1*time.Hour)
	return
}
