package main

import (
	"fmt"
	"time"
)

func combine(inCh1, inCh2 <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				return
			}
		}()
		defer close(out)

		for {
			select {
			case r, ok := <-inCh1:
				if !ok {
					inCh1 = nil
					break
				}
				out <- r
			case r, ok := <-inCh2:
				if !ok {
					inCh2 = nil
					break
				}
				out <- r
			}

			if inCh1 == nil && inCh2 == nil {
				break
			}
		}
	}()

	return out
}

func main() {
	ch1 := make(<-chan int)
	ch2 := make(<-chan int)
	outCh := make(<-chan int)
	outCh = combine(ch1, ch2)

	time.Sleep(time.Second * 5)
	for msg := range outCh {
		fmt.Println(msg)
	}
}
