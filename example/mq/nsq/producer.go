package main

import (
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("119.91.202.245:4150", config)
	if err != nil {
		panic(err)
	}

	topicName := "test_topic_1"
	for {
		msgBody := []byte("message=" + time.Now().String())
		err := producer.Publish(topicName, msgBody)
		if err != nil {
			log.Printf("推送失败：%v", err)
		}
		log.Println("Sent successfully")
		time.Sleep(2e9)
	}

	//sigChan := make(chan os.Signal, 1)
	//signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	//<-sigChan

	producer.Stop()

}
