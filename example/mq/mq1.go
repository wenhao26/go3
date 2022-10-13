package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/streadway/amqp"

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
	err := mq.Channel.ExchangeDeclare(
		"test.20220820",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(fmt.Errorf("声明交换机失败: %s \n", err))
	}
	defer mq.Channel.Close()

	for {
		data := map[string]string{
			"time": time.Now().Format("2006-01-02 15:04:05"),
		}
		message, _ := json.Marshal(data)
		err = mq.Channel.Publish(
			"test.20220820",
			"test.20220820_key",
			false,
			false,
			amqp.Publishing{
				// DeliveryMode: amqp.Persistent, // 消息持久化
				ContentType: "text/plain",
				Body:        message,
			})
		if err != nil {
			fmt.Println("推送消息异常：", err)
		}
		fmt.Printf(" [x] Sent %s\n", message)
		time.Sleep(time.Millisecond * 100)
	}
}
