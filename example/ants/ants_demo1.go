package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

func Task() {
	fmt.Println("output。。。")
	time.Sleep(time.Second * 1)
}

func main() {
	p, _ := ants.NewPool(10)
	defer p.Release()

	var wg sync.WaitGroup
	for {
		if p.Running() < 10 {
			wg.Add(1)
			p.Submit(func() {
				// TODO
				Task()
				wg.Done()
			})
		}
	}
	wg.Wait()
}
