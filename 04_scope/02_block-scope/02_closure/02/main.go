package main

import (
	"fmt"
)

var x = 0 // package level declaration influenced by inner scopes

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
}
