package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func GetFiles(path string) {
	fs, _ := ioutil.ReadDir(path)
	for _, file := range fs {
		wg1.Add(1)
		go func(file os.FileInfo) {
			defer wg1.Done()
			if file.IsDir() {
				fmt.Println(path + file.Name() + " 属于文件夹...")
				GetFiles(path + file.Name() + "/")
			} else {
				fmt.Println(" --" + path + file.Name())
			}
		}(file)
	}
}

func GetFiles2(path string) {
	fs, _ := ioutil.ReadDir(path)
	for _, file := range fs {
		if file.IsDir() {
			fmt.Println(path + file.Name() + " 属于文件夹...")
			GetFiles(path + file.Name() + "/")
		} else {
			fmt.Println(" --" + path + file.Name())
		}
	}
}

func GetFiles3(path string) {
	fs, _ := ioutil.ReadDir(path)
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		for _, file := range fs {
			if file.IsDir() {
				fmt.Println(path + file.Name() + " 属于文件夹...")
				GetFiles(path + file.Name() + "/")
			} else {
				fmt.Println(" --" + path + file.Name())
			}
		}
	}()
}

func main() {
	start := time.Now()
	path := "D:\\www\\"
	GetFiles3(path)
	wg1.Wait()
	//GetFiles2(path)
	elapsed := time.Since(start)
	fmt.Println("完成耗时：", elapsed)
}
