package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		panic(err)
	}

	topicName := "test_topic_1"
	go func() {
		for {
			msgBody := []byte("message=" + time.Now().String())
			err := producer.Publish(topicName, msgBody)
			if err != nil {
				log.Printf("推送失败：%v", err)
			}
			log.Println("Sent successfully")
			// time.Sleep(2e9)
			time.Sleep(time.Millisecond * 100)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	producer.Stop()
}
