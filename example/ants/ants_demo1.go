package main

import (
	"fmt"

	"github.com/panjf2000/ants/v2"
)

func main() {
	p, _ := ants.NewPool(10)

	fmt.Println(p)
}
