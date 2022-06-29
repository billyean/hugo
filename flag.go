package main

import (
	"flag"
	"fmt"
	"strings"
)

type ArrayValue []string

func (s *ArrayValue) String() string {
	return fmt.Sprintf("%v", *s)
}

func (a *ArrayValue) Set(s string) error {
	*a = strings.Split(s, ",")
	return nil
}

// Diff with args with unix style --
// -member 5 -names Haibo,Tina,Ruby,Tristan,Zachary
// We got Haibo
// We got Tina
// We got Ruby
// We got Tristan
// We got Zachary
func main() {
	member := flag.Int("member", 0, "Defines family number.")

	var names ArrayValue
	flag.Var(&names, "names", "Input family member's name")
	flag.Parse()

	for i := 0; i < *member; i++ {
		fmt.Printf("We got %s\n", names[i])
	}
}
