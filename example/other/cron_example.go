package main

import (
	"log"
	"time"

	"github.com/robfig/cron"
)

func DelTag() {
	log.Println("Deleted!")
}

func main() {
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("*/5 * * * * *", func() {
		log.Println("Run DelTag...")
		DelTag()
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
