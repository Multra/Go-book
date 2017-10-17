package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	for v, arg := range os.Args[1:] {
		fmt.Println(v, arg)
	}
	fmt.Println(s)
}
