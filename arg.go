package main

import (
	"fmt"
	"os"
)

func main() {
	// Accessing program arguments
	args := os.Args
	fmt.Println(args[1:])
}
