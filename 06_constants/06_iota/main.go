package main

import (
	"fmt"
)

const (
	_ = iota // 0
	// B comment to make Visual Studio Code happy
	B = iota * 10 // 1 * 10
	//C comment to make Visual Studio Code happy
	C = iota * 10 // 2 * 10
)

func main() {
	fmt.Println(B)
	fmt.Println(C)
}
