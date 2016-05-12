package main

import (
	"fmt"
)

const (
	message = "你好，%s, %d!\n"
	name    = "严海波"
	name2   = iota
	name3
)

func main() {
	b := 65
	fmt.Printf("it's %c\n", b)
}
