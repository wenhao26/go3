package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/panjf2000/ants/v2"
	"github.com/streadway/amqp"

	"coinsky_go_project/example/utils"
)

func Consumer(conn *amqp.Connection, queueName string) {
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Fatal error Channel: %s \n", err)
	}

	/*ch.Qos(
		1000,
		0,
		false,
	)*/
	defer ch.Close()

	messages, err := ch.Consume(
		queueName,
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
		log.Printf("-Msg=%s", string(message.Body))
		// _ = message.Ack(true) // Ack
	}
}

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

	// 队列名称
	queueName := viper.GetString("rabbitmq.test_queue")

	// 使用线程池
	var wg sync.WaitGroup
	runTimes := 20
	p, _ := ants.NewPool(runTimes, ants.WithPreAlloc(false))
	defer p.Release()
	for {
		if p.Running() < runTimes {
			wg.Add(1)
			p.Submit(func() {
				Consumer(conn, queueName)
				wg.Done()
			})
			fmt.Println("开启运行:", p.Running())
		}
	}
	wg.Wait()
}
