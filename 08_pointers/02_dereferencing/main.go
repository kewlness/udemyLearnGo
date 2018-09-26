package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // memory location

	var b = &a      // var b *int = &a
	fmt.Println(b)  // same address as &a
	fmt.Println(*b) // 43 --> dereferencing b

	// b is an int pointer
	// b points to the memory address where an int is stored
	// to see the value in that memory address, add a * in front of b
	// this is known as dereferencing
	// the * is an operator in this case
}
