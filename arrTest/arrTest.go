package main

import (
	"fmt"
)

// Note Go is pass by value, it doesn't change original
// object.
func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s,", word)
	}
	fmt.Printf("\n")
	w[2] = "blue"
}

// LeetCode's exercise. Note this doesn't have boundary checking.
// https://leetcode.com/problems/counting-bits/
func count(n int) []int {
	var bits []int

	// first value is 0
	bits = append(bits, 0)
	v := 1
	var p uint = 0

	for v <= n {
		start := (1 << p)
		end := start * 2
		for v < end && v <= n {
			bits = append(bits, 1 + bits[v - start])
			v = v + 1
		}
		p = p + 1
	}

	return bits
}

func main() {
	//words := []string{"the", "quick", "brown", "fox"}
	//printer(words[1:])
	//printer(words)
	//
	//word1 := make([]string, 4)
	//word1[0] = "hello"
	//word1[1] = "world"
	//word1[2] = "good"
	//word1[3] = "morning"
	//printer(word1)
	//word1 = append(word1, "go")
	//printer(word1)
	//
	//word2 := make([]string, 4)
	//copy(word2, word1)
	//printer(word2)
	//word2[1] = "你好"
	//printer(word1)
	//printer(word2)
	bitsCount := count(17)
	fmt.Printf("[")

	i := 0

	for i < len(bitsCount) {
		fmt.Printf(" %d ", bitsCount[i])
		i = i + 1
	}
	fmt.Printf("]\n")
}
