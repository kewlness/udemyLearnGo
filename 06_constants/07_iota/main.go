package main

import (
	"fmt"
)

const (
	_ = iota //0
	// KB is kilobytes (yes, a comment to make Visual Studio Code happy)
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	// MB is megabytes (and yet another comment to make Visual Studio Code happy)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
)

func main() {
	fmt.Println("Binary\t\tDecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
}
