package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var rw ReadWriter

	// os.Stdin implements Reader
	rw = os.Stdin
	n, err := fmt.Fscanln(rw);
	if (err == nil) {

		// os.Stdout implements Writer
		rw = os.Stdout
		fmt.Fprintf(rw, "hello, writer %d \n", n)
	} else {

		// os.Stdout implements Writer
		rw = os.Stdout
		fmt.Fprintf(rw, "something went wrong %s \n", err)
	}

}
