package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	fmt.Println("len: ", len(s))
	n := 0
	for _, _ = range s {
		n++
	}
	fmt.Println("rune count:", n)

	n = 0
	for range s {
		n++
	}
	fmt.Println("rune count:", n)

	fmt.Println("rune count:", utf8.RuneCountInString(s))
	r := []rune("ë")
	fmt.Println(string(unicode.ToUpper(r[0])))

	fmt.Println(strings.ToUpper("ë"))
}
