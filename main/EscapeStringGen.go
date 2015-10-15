package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {

	file, err := os.Open("escapestring.file")
	if err == nil {
		fmt.Println(file)

		var ret []byte
		buf := make([]byte, 512)

		for {
			n, err := file.Read(buf)
			if err == nil && n > 0 {
				fmt.Println(n)
				fmt.Println(buf)
				ret = append(ret, buf[:n]...)
			} else {
				fmt.Println(err)
				break
			}
		}

		fmt.Println(strings.Replace(string(ret), "\"", "\\\"", -1))
	}
}
