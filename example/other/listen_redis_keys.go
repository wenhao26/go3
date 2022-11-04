package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var wg sync.WaitGroup
var ctx = context.Background()
var redisCli *redis.Client

func init() {
	redisCli = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
	})
}

// 订阅key过期事件
func SubExpireEvent() {
	//sub := redisCli.Subscribe(ctx, "__keyevent@0__:expired")
	sub := redisCli.Subscribe(ctx, "__key*@0__:*")

	// 这里通过一个for循环监听redis-server发来的消息。
	// 当客户端接收到redis-server发送的事件通知时，
	// 客户端会通过一个channel告知我们。我们再根据
	// msg的channel字段来判断是不是我们期望收到的消息，
	// 然后再进行业务处理。
	for {
		msg := <-sub.Channel()
		log.Println("Channel:", msg.Channel)
		log.Println("Pattern :", msg.Pattern)
		log.Println("Payload:", msg.Payload)
		log.Println("PayloadSlice :", msg.PayloadSlice)
		fmt.Println("\n")
	}
}

func main() {
	fmt.Println("订阅键名过期事件...")

	redisCli.Set(ctx, "test", "test", 8*time.Second)

	wg.Add(1)
	go SubExpireEvent()
	defer wg.Done()
	wg.Wait()
}
