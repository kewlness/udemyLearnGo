package main

import (
	"fmt"
)

const (
	// A comment to make Visual Studio Code happy
	A = iota //0
	// B comment to make Visual Studio Code happy
	B
	// C comment to make Visual Studio Code happy
	C
)

const (
	// D comment to make Visual Studio Code happy
	D = iota //0
	// E comment to make Visual Studio Code happy
	E //1
	// F comment to make Visual Studio Code happy
	F //2
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
}
