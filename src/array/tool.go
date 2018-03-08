package array

import (
	"fmt"
)

// Note Go is pass by value, it doesn't change original
// object.
func Printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s,", word)
	}
	fmt.Printf("\n")
	w[2] = "blue"
}

// LeetCode's exercise. Note this doesn't have boundary checking.
// https://leetcode.com/problems/counting-bits/
func Count(n int) []int {
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

