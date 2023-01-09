package main

import (
	"fmt"
)

func Counter1(in chan<- int) {
	defer close(in)
	for i := 1; i <= 1000; i++ {
		in <- i
	}
}

func Counter2(in chan<- int, out <-chan int) {
	defer close(in)
	for i := range out {
		in <- i + i
	}
}

func Output(out <-chan int) {
	for i := range out {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Counter1(ch1)
	go Counter2(ch2, ch1)
	Output(ch2)

}
