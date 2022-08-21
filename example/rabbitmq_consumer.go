package main

import (
	"coinsky_go_project/example/utils"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	viper := utils.LoadConfig()

	conn, err := amqp.Dial(fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		viper.GetString("rabbitmq.username"),
		viper.GetString("rabbitmq.password"),
		viper.GetString("rabbitmq.host"),
		viper.GetInt("rabbitmq.port")),
	)
	if err != nil {
		panic(fmt.Errorf("Fatal error Connection: %s \n", err))
	}
	defer conn.Close()

	forever := make(chan bool)
	for i := 1; i <= 6; i++ {
		go func(routineNum int) {
			ch, err := conn.Channel()
			if err != nil {
				log.Printf("Fatal error Channel: %s \n", err)
			}
			defer ch.Close()

			messages, err := ch.Consume(
				viper.GetString("rabbitmq.test_queue"),
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
				log.Printf("-routineNum#%d：%s", routineNum, string(message.Body))
				// _ = message.Ack(true) // Ack
			}

		}(i)
	}
	fmt.Println("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
