package main

import (
	"fmt"
)

func main() {
	//functions below have identical behaviors
	newInt()
	newInt2()

	p := new(int)
	q := new(int)
	fmt.Println(p == q) //"false"
}

func newInt() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}
