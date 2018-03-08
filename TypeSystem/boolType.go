package main

import "fmt"

func main() {
	var message = "Hello Tristan...\t"
	a, b, c := true,20, true
	fmt.Print(message, a, b, c)

	var say *string = &message
	fmt.Print("\n", *say)
}
