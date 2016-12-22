package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type IntSet struct { /***/
}

func (*IntSet) String() string

func main() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	//w = time.Second // compile error: time.Duration lacks Write method

	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	// rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method

	w = rwc
	// rwc = w // compile error: io.Writer lacks Close method

	// var _ = IntSet{}.String() // compile error: String requires *IntSet receiver
	var s IntSet
	var _ = s.String() // OK: s is a variable and &s has a String method

	var _ fmt.Stringer = &s
	// var _ fmt.Stringer = s // compile error: IntSet lacks String method

	// empty interface
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)

	//*bytes.Buffer must satisfy io.Writer
	var w1 io.Writer = new(bytes.Buffer)
	var _ io.Writer = (*bytes.Buffer)(nil)

}
