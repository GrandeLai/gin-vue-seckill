package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@192.168.10.3:5672")
	fmt.Println(err)
	defer conn.Close()

	ch, err1 := conn.Channel()
	fmt.Println(err1)
	defer ch.Close()

	ch.QueueDeclare("first_queue", true, false, false, false, nil)
	ch.QueueDeclare("second_queue", true, false, false, false, nil)

	//创建交换机，direct直连类型交换机
	ch.ExchangeDeclare("first_exchange", "direct", true, false, false, false, nil)
	ch.ExchangeDeclare("second_exchange", "direct", true, false, false, false, nil)

	//交换机绑定
	ch.QueueBind("first_queue", "first_routingKey", "first_exchange", false, nil)
	ch.QueueBind("second_queue", "second_routingKey", "second_exchange", false, nil)

	//发布任务
	ch.Publish("first_exchange", "first_routingKey", false, false, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         []byte("hello first_exchange"),
		DeliveryMode: amqp.Persistent,
	}) //direct和routingkey要对应起来
	ch.Publish("second_exchange", "second_routingKey", false, false, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         []byte("hello second_exchange"),
		DeliveryMode: amqp.Persistent,
	})
}
