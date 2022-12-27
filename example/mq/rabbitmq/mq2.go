package main

import (
	"log"

	"coinsky_go_project/common/rabbitmq"
)

func main() {
	opt := rabbitmq.Option{
		Username: "admin",
		Password: "admin",
		Host:     "192.168.1.202",
		Port:     5672,
	}
	mq := rabbitmq.NewConn(opt)
	defer mq.Connection.Close()

	forever := make(chan bool)
	go func() {
		//mq.Channel.Qos(1000, 0, false)
		messages, err := mq.Channel.Consume(
			"test.queue_01",
			"",
			true, // 自动确认
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Println("消费消息异常：", err)
		}

		for message := range messages {
			log.Printf("-MSG=%s", string(message.Body))
			// _ = message.Ack(true) // Ack
		}
	}()
	<-forever
}
