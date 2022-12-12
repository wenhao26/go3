package main

import (
	"fmt"
)

func main() {
	var documents []map[string]interface{}

	documents = append(documents, map[string]interface{}{
		"id": 1, "title": "测试一下1",
	})
	documents = append(documents, map[string]interface{}{
		"id": 2, "title": "测试一下2",
	})

	fmt.Println(documents)
}
