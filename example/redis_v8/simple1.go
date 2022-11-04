package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Rdb struct {
	Client *redis.Client
}

func RedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}

func (r *Rdb) Set(key string, value interface{}, ttl time.Duration) {
	err := r.Client.Set(ctx, key, value, ttl).Err()
	if err != nil {
		log.Fatal(err)
	}
}

func (r *Rdb) Get(key string) string {
	ret, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}

	return ret
}

func main() {
	key := "test:string_key"
	value := "测试一下"
	ttl := 60 * time.Second

	rdb := Rdb{Client: RedisClient()}
	rdb.Set(key, value, ttl)

	fmt.Println("Done...")
}
