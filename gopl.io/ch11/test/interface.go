package main

import (
	"io"
	//`    "os"
	"bytes"
	"fmt"
)

func main() {
	//    var w io.Writer
	var buf *bytes.Buffer
	//w = os.Stdout
	//w = new(bytes.Buffer)
	//w = nil
	if buf == nil {
		fmt.Println("It's nil, man")
	}
	f(buf)
}

func f(w io.Writer) {
	if w != nil {
		fmt.Println("It's nil, man")
	}
}
