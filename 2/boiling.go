package main

import "fmt"

const boilingF = 212.0

func main() {
	var c = (boilingF - 32) * 5 / 9
	fmt.Printf("boiling point = %v F or %v C\n", boilingF, c)
}
