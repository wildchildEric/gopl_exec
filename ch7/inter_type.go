package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	w = nil

	var x interface{} = time.Now()

	var x1 interface{} = []int{1, 2, 3}
	fmt.Println(x1 == x1) // panic: comparing uncomparable type []int

	var w1 io.Writer
	fmt.Printf("%T\n", w1) // "<nil>"

	w1 = os.Stdout
	fmt.Printf("%T\n", w1) // "*os.File"

	w1 = new(bytes.Buffer)
	fmt.Printf("%T\n", w1) // "*bytes.Buffer"
}
