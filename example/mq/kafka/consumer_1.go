package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

var wg sync.WaitGroup

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_1_1_0
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = time.Second * 5
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	client, err := sarama.NewClient([]string{"192.168.1.216:9092", "192.168.1.217:9092", "192.168.1.218:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	topic := "topic名称"
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println("无法获取分区列表：", err)
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("无法启动分区[%d]的使用者：%s\n", partition, err)
			return
		}
		defer pc.AsyncClose()

		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()

			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)

		wg.Wait()
		consumer.Close()
	}

}
