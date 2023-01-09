package main

import (
	"fmt"
)

func ArrayColumn(input map[string]map[string]interface{}, columnKey string) []interface{} {
	columns := make([]interface{}, 0, len(input))
	for _, val := range input {
		if v, ok := val[columnKey]; ok {
			columns = append(columns, v)
		}
	}
	return columns
}

func main() {
	input := make(map[string]map[string]interface{})

	input["1"] = map[string]interface{}{
		"name": "Testing1",
	}
	input["2"] = map[string]interface{}{
		"name": "Testing2",
	}
	input["3"] = map[string]interface{}{
		"name": "Testing3",
	}
	ret := ArrayColumn(input, "name")
	fmt.Println(ret)
}
