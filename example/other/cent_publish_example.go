package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/centrifugal/gocent/v3"
	"github.com/robfig/cron"
	"github.com/shirou/gopsutil/mem"
)

type Data struct {
	Info string `json:"info"`
	Mem  string `json:"mem"`
}

func Publish(cent gocent.Client) {
	data := Data{
		Info: time.Now().Format("2006-01-02 15:04:05"),
		Mem:  getMemInfo(),
	}
	pushData, _ := json.Marshal(data)

	/*cent := gocent.New(gocent.Config{
		Addr: "http://119.91.202.245:9501/api",
		Key:  "7c2da7dd-c51f-430b-966f-e39bb4254346",
	})*/

	ch := "channel"
	ctx := context.Background()

	//result, err := cent.Publish(ctx, ch, []byte(`{"test":"测试一下..."}`))
	result, err := cent.Publish(ctx, ch, pushData)
	if err != nil {
		log.Fatalf("调用发布时出错: %v", err)
	}
	log.Printf("发布到频道 %s 成功, 流位置 {offset: %d, epoch: %s}", ch, result.Offset, result.Epoch)
}

func getMemInfo() string {
	info, _ := mem.VirtualMemory()
	return info.String()
}

func main() {
	cent := gocent.New(gocent.Config{
		Addr: "http://119.91.202.245:9501/api",
		Key:  "7c2da7dd-c51f-430b-966f-e39bb4254346",
	})

	c := cron.New()
	_ = c.AddFunc("*/5 * * * * *", func() {
		log.Println("Run Publish...")
		Publish(*cent)
	})
	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
