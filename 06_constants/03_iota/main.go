package main

import (
	"fmt"
)

const (
	// A comment to make Visual Studio Code happy
	A = iota //0
	// B comment to make Visual Studio Code happy
	B = iota //1
	// C comment to make Visual Studio Code happy
	C = iota //2
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

}
