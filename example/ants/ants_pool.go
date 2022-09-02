package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

func demoFunc2() {
	time.Sleep(2 * time.Second)
	fmt.Println("ant test。。。")
}

func main() {
	// 释放ants的默认协程池
	defer ants.Release()

	var wg sync.WaitGroup

	// 任务函数
	syncCalculateSum := func() {
		demoFunc2()
		wg.Done()
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)

		// 提交任务到默认协程池
		ants.Submit(syncCalculateSum)
	}

	wg.Wait()

	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")
}
