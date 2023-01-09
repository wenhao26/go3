package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Pub(msg chan<- string) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				return
			}
		}()

		for {
			dateStr := time.Now().String()
			msg <- dateStr
			fmt.Println("Sent:", dateStr)
			time.Sleep(time.Millisecond * 100)
		}
	}()
}

func Sub(msg <-chan string) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				return
			}
		}()

		for {
			select {
			case m, ok := <-msg:
				if !ok {
					return
				}
				fmt.Println("Receive:", m)
			}
		}
	}()
}

func main() {
	msg := make(chan string, 10)
	defer close(msg)

	Sub(msg)
	Pub(msg)

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	<-c
}
