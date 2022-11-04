package main

import (
	"fmt"
	"time"
)

func main() {
	// #1
	/*c := make(chan int)
	o := make(chan bool)

	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o*/

	// #2
	/*c := make(chan int)
	go func() {
		c <- 1
		time.Sleep(time.Second)
		c <- 2
		fmt.Println("send complete！")
	}()

	fmt.Println("主线程")
	time.Sleep(time.Second)
	i := <- c
	fmt.Printf("receive %d\n", i)
	i = <- c
	fmt.Printf("receive %d\n", i)
	time.Sleep(time.Second)*/

	// #3
	c := make(chan int, 2)
	go func() {
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Printf("send %d\n", i)
		}
		time.Sleep(3 * time.Second)
		for i := 3; i < 5; i++ {
			c <- i
			fmt.Printf("send %d\n", i)
		}
	}()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Printf("receive %d\n", <-c)
	}

	fmt.Println("Done")

}
