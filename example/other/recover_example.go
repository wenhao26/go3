package main

import (
	"fmt"
)

func TestFunc() {
	fmt.Println("TestFunc")
	panic("异常")
}

func main() {
	fmt.Println("Start")
	defer func() {
		fmt.Println("----------")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("**********")
	}()
	TestFunc()
	fmt.Println("End")
}
