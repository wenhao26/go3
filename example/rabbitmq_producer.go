package main

import (
	"coinsky_go_project/example/utils"
	"fmt"
	"github.com/streadway/amqp"
	"time"
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

	ch, err := conn.Channel()
	if err != nil {
		panic(fmt.Errorf("Fatal error Channel: %s \n", err))
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		viper.GetString("rabbitmq.test_exchange"),
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(fmt.Errorf("Fatal error Exchange: %s \n", err))
	}

	forever := make(chan bool)
	go func() {
		for {
			message := "time=" + time.Now().Format("2006-01-02 15:04:05")
			err = ch.Publish(
				viper.GetString("rabbitmq.test_exchange"),
				viper.GetString("rabbitmq.test_routing_key"),
				false,
				false,
				amqp.Publishing{
					// DeliveryMode: amqp.Persistent, // 消息持久化
					ContentType: "text/plain",
					Body:        []byte(message),
				})
			if err != nil {
				fmt.Println("推送消息异常：", err)
			}
			fmt.Printf(" [x] Sent %s\n", message)
		}
	}()
	fmt.Println("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
