package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

// 消费者类型
type ConsumerHandler struct {
	Title string
}

// 实现处理消息的方法
func (c *ConsumerHandler) HandleMessage(msg *nsq.Message) error {
	if len(msg.Body) == 0 {
		fmt.Println("无消息...")
		return nil
	}
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

func main() {
	topicName := "test_topic_1"
	channel := "test_channel_1"

	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	consumer, err := nsq.NewConsumer(topicName, channel, config)
	if err != nil {
		panic(err)
	}

	// consumer.SetLogger(nil, 0)
	consumer.AddHandler(&ConsumerHandler{
		Title: "消费者1",
	})

	if err := consumer.ConnectToNSQLookupd("127.0.0.1:4161"); err != nil {
		panic(err)
	}

	/*sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan*/
	c := make(chan os.Signal)        // 定义一个信号的通道
	signal.Notify(c, syscall.SIGINT) // 转发键盘中断信号到c
	<-c                              // 阻塞

	consumer.Stop() // 优雅地停止消费者
}
