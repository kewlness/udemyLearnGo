package main

import (
	"fmt"
)

func main() {
	x := 42 // outer scope
	fmt.Println(x)
	{ // random curly braces to provide an inner scope to func main
		fmt.Println(x) // accesible to inner scope
		y := "The credit belongs with the one who is in the ring"
		fmt.Println(y) // defined in inner scope so accessible
	}
	// fmt.Println(y) // not accessible because not defined in outer scope
}
