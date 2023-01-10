package main

import (
	"fmt"
	"math/rand"
	"time"
)

func eat() chan string {
	out := make(chan string)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				return
			}
		}()

		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		out <- "talking to the moon"
		close(out)
	}()
	return out
}

func main() {
	// CASE1
	/*readCh := make(chan int, 1)
	writeCh := make(chan int, 1)

	num := 100
	select {
	case ret := <-readCh:
		fmt.Println("读通道：", ret)
	case writeCh <- num:
		fmt.Println("写通道", num)
	default:
		fmt.Println("TODO。。。")
	}*/

	// CASE2
	eatCh := eat()
	sleep := time.NewTimer(time.Second * 1)
	select {
	case r := <-eatCh:
		fmt.Println(r)
	case <-sleep.C:
		fmt.Println("Time to sleep")
	default:
		fmt.Println("Fly。。。")
	}

}
