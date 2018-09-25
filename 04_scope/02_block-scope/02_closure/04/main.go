package main

import (
	"fmt"
)

func wrapper() func() int { // return type func() int - outer scope
	x := 0
	return func() int { // this is the func being returned - inner scope
		x++
		return x // value being returned by the func
	}
}

func main() {
	increment := wrapper() // func expression
	fmt.Println(increment())
	fmt.Println(increment())
}
