package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

func main() {
	var buf *bytes.Buffer
	// var buf io.Writer //To Fix the panic
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf) // NOTE: subtly incorrect!
	if debug {
		//...use buf...
	}
}

func f(out io.Writer) {
	//... do someting ...
	fmt.Printf("%[1]T %[1]v\n", out)
	if out != nil {
		out.Write([]byte("done!\n")) // panic: nil pointer dereference
	}
}
