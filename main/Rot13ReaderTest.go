package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

const (
	A = byte('A')
	M = byte('M')
	Z = byte('Z')
	a = byte('a')
	m = byte('m')
	z = byte('z')
)

func (r *rot13Reader) Read(bs []byte) (int, error) {
	n, err := r.r.Read(bs)
	if err != nil {
		return n, err
	}
	for i, v := range bs[:n] {
		if v >= A && v <= Z {
			if v <= M {
				bs[i] = v + 13
			} else {
				bs[i] = v - 13
			}
		}
		if v >= a && v <= z {
			if v <= m {
				bs[i] = v + 13
			} else {
				bs[i] = v - 13
			}
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
