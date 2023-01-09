package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"

	"github.com/Jeffail/tunny"
)

const (
	DataSize = 10000
	DataTask = 100
)

func main() {
	numCPUs := runtime.NumCPU()
	pool := tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {
		var sum int
		for _, n := range payload.([]int) {
			sum += n
		}
		return sum
	})
	defer pool.Close()

	// 平均随机分组
	nums := make([]int, DataSize)
	for i := range nums {
		nums[i] = rand.Intn(1000)
	}

	var wg sync.WaitGroup
	wg.Add(DataSize / DataTask)
	parSums := make([]int, DataSize/DataTask)
	for i := 0; i < DataSize/DataTask; i++ {
		go func(i int) {
			parSums[i] = pool.Process(nums[i*DataTask : (i+1)*DataTask]).(int)
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println(parSums)

}
