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

	//声明队列
	queue, err2 := ch.QueueDeclare("my queue", false, false, false, false, nil)
	fmt.Println(err2)
	fmt.Println(queue)

	err_p := ch.Publish("", queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("hello world"),
	})
	fmt.Println(err_p)

	// 消费者
	msgs, err_c := ch.Consume("my queue", "my_consumer", false, false, false, false, nil)
	fmt.Println(err_c)

	for msg := range msgs { // chan类型
		// DeliveryTag:唯一标识
		fmt.Println(msg.DeliveryTag, string(msg.Body))
	}
}
