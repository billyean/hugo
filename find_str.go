package main

import (
	"fmt"
	"strings"
)
const target = "stay foolish stay hungry"

func main() {
	foolish_str_str := "foolish"
	contains_stay := strings.Contains(target, foolish_str_str)
	fmt.Printf("%s contains %s: %t\n", target, foolish_str_str, contains_stay)

	prefix_str := "stay"
	prefix_t := strings.HasPrefix(target, prefix_str)
	fmt.Printf("%s has prefix %s: %t\n", target, prefix_str, prefix_t)

	suffix_str := "stay"
	suffix_t := strings.HasSuffix(target, suffix_str)
	fmt.Printf("%s has psuffix %s: %t\n", target, suffix_str, suffix_t)

	// Split
	words := strings.Split(target, " ")
	for _, word := range words {
		fmt.Println(word)
	}
}
