package main

import (
	"fmt"
	"strings"
)

func Explode(delimiter, text string) []string {
	return strings.Split(text, delimiter)
}

func main() {
	text := "a,b,c,d"
	arr := Explode(",", text)
	fmt.Println(arr)
}
