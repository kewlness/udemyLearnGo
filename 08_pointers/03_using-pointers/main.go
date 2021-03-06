package main

import (
	"fmt"
)

func main() {
	a := 43
	fmt.Println(a)  // 43
	fmt.Println(&a) // memory address of a

	var b = &a      // var b *int = &a
	fmt.Println(b)  // memory address of a
	fmt.Println(*b) // 43

	*b = 42        // Change the value at this address to 42
	fmt.Println(a) // 42

	// this is useful
	// we can pass a memory address instead of a bunch of values (we can pass a reference)
	// and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// we don't have to pass around large amounts of data
	// we only have to pass around addresses

	// everything is PASS BY VALUE in go, btw
	// when we pass a memory address, we are passing a value
}
