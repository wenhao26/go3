package main

import (
	"fmt"
	"sync"
)

// 启动一个协程，生产任务，写入到jobCh
func genJob1(n int) <-chan int {
	jobCh := make(chan int, 200)
	go func() {
		for i := 0; i < n; i++ {
			jobCh <- i
		}
		close(jobCh)
	}()
	return jobCh
}

var wg3 *sync.WaitGroup = &sync.WaitGroup{}

func workerPool1(n int, jobCh <-chan int, retCh chan<- string) {
	for i := 0; i < n; i++ {
		wg3.Add(1)
		go worker1(i, jobCh, retCh)
	}
}

func worker1(tid int, jobCh <-chan int, retCh chan<- string) {
	cnt := 0
	for job := range jobCh {
		cnt++
		ret := fmt.Sprintf("worker %d processed job: %d, it's the %dth processed by me.", tid, job, cnt)
		retCh <- ret
	}
	wg3.Done()
}

func main() {
	jobCh := genJob1(10000)
	retCh := make(chan string, 10000)
	workerPool1(5, jobCh, retCh)

	/*time.Sleep(time.Second * 1)
	  close(retCh)
	  for ret := range retCh {
	    fmt.Println(ret)
	  }*/

	// 先创建1000个任务模型（goroutine）
	// 通过协程池子的方式，去同步读取任务模型，并写入结果通道
	// 读取结果通道结果
	wg3.Wait()
	close(retCh)
	for {
		select {
		case r, ok := <-retCh:
			if !ok {
				fmt.Println("done")
				return
			}
			fmt.Println("job =>", r)
		}
	}

}

