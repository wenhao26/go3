package main

import (
	"fmt"
)

func ArrayUnique(arr []string) []string {
	size := len(arr)
	result := make([]string, 0, size)
	temp := map[string]interface{}{}
	for i := 0; i < size; i++ {
		if _, ok := temp[arr[i]]; ok != true {
			temp[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}
	return result
}

func main() {
	arr := []string{"a", "a", "b", "c", "c", "d"}
	ret := ArrayUnique(arr)
	fmt.Println(ret)
}
