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

	deleveries, err := ch.Consume("second_queue", "second_consumer", false, false, false, false, nil)
	if err == nil {
		for delever := range deleveries {
			fmt.Println(string(delever.Body))
			delever.Ack(true)
		}
	}
}
