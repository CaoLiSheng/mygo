package main

import (
	"fmt"
)

type MyUint struct {
	value uint
}

func main()  {
	fmt.Println("hello go!")

	pI := &MyUint{1}

	fmt.Println(pI)
	fmt.Println(*pI)
}
