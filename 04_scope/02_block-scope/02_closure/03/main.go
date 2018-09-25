package main

import (
	"fmt"
)

func main() {
	x := 0                    // outer scope
	increment := func() int { // inner scope
		x++
		return x
	} // end inner scope
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope

increment := func() int{...}
anonymous function
a function without a name
basically a lambda in python

func expression
assigning a func to a variable
*/
