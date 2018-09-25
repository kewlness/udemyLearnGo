package main

import (
	"fmt"
)

func main() {
	// will not compile
	// order matters
	fmt.Println(x)
	fmt.Println(y)
	x := 42 // scope of x begins here and goes to the closing brace
}

var y = 42 // scope of y begins here and goes to the end of the program
