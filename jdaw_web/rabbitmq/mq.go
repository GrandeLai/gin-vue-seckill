package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

type RabbitMq struct {
	Conn         *amqp.Connection
	Ch           *amqp.Channel
	QueueName    string
	ExchangeName string
	ExchangeType string
	RoutingKey   string
}

type QueueAndExchange struct {
	QueueName    string
	ExchangeName string
	ExchangeType string
	RoutingKey   string
}

func NewRabbitMq(qe QueueAndExchange) RabbitMq {
	return RabbitMq{
		QueueName:    qe.QueueName,
		ExchangeName: qe.ExchangeName,
		ExchangeType: qe.ExchangeType,
		RoutingKey:   qe.RoutingKey,
	}
}

func (r *RabbitMq) ConnMq() {
	conn, err := amqp.Dial("amqp://admin:admin@192.168.10.3:5672")
	if err != nil {
		fmt.Printf("连接mq出错，错误信息：%v\n", err)
		return
	}
	r.Conn = conn

	ch, err1 := conn.Channel()
	fmt.Println(err1)
	defer ch.Close()
}

func (r *RabbitMq) CloseMq() {
	err := r.Conn.Close()
	if err != nil {
		fmt.Printf("关闭mq连接出错，错误信息：%v\n", err)
		return
	}
}

func (r *RabbitMq) OpenChannel() {
	ch, err := r.Conn.Channel()
	if err != nil {
		fmt.Printf("开启channel通道出错，错误信息：%v\n", err)
		return
	}
	r.Ch = ch
}

func (r *RabbitMq) CloseChannel() {
	err := r.Ch.Close()
	if err != nil {
		fmt.Printf("关闭channel通道出错，错误信息：%v\n", err)
		return
	}
}

func (r *RabbitMq) PublishMsg(body string) {
	ch := r.Ch
	ch.QueueDeclare(r.QueueName, true, false, false, false, nil)
	ch.ExchangeDeclare(r.ExchangeName, r.ExchangeType, true, false, false, false, nil)
	ch.QueueBind(r.QueueName, r.RoutingKey, r.ExchangeName, false, nil)
	ch.Publish(r.ExchangeName, r.RoutingKey, false, false, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         []byte(body),
		DeliveryMode: amqp.Persistent,
	})
}
