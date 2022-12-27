package main

import (
	"fmt"

	"github.com/uniplaces/carbon"
)

func main() {
	fmt.Println(carbon.Now().DateTimeString())
}
