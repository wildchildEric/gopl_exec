package main

import (
	"fmt"
)

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var a1 [3]int = [3]int{1, 2, 3}
	for i, v := range a1 {
		fmt.Printf("%d %d\n", i, v)
	}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2])

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])

	r1 := [...]int{99: -1}
	fmt.Printf("%v\n", r1)

	a2 := [2]int{1, 2}
	b2 := [...]int{1, 2}
	c2 := [2]int{1, 3}
	fmt.Println(a2 == b2, a2 == c2, b2 == c2)
	// invalid operation: a2 == d2 (mismatched types [2]int and [3]int)
	// d2 := [3]int{1, 2}
	// fmt.Println(a2 == d2)

}
