package main

import "fmt"

func main() {
	defer fmt.Println("world1")

	{
		defer fmt.Println("world2")

		{
			defer fmt.Println("world3")

			fmt.Println("hello")
		}
	}
}