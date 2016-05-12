package main

import (
	"fmt"
)

// I am extremely no happy that abs(int) is not a supported
// function in Go, I also noticed people are talking it's not
// worth to be added since it's too generic.
// This won't make GO a better language than C++ because we need
// re-invent the wheel all the time.
func abs(v int) int {
	if v >= 0 {
		return v
	} else {
		return -v
	}
}

// LeetCode exercise : power of 4 is a trick function
// https://leetcode.com/problems/power-of-four/
func isPowerOfFour(v int) string{
	v1 := abs(v)
	yes := v1 & (v1 - 1) == 0 && (v1 - 1) % 3 == 0
	yes_or_not := ""
	if !yes {
		yes_or_not = "not "
	}
	return yes_or_not
}

func main() {
	fmt.Printf("%d is %spower of four\n", 5, isPowerOfFour(5))
	fmt.Printf("%d is %spower of four\n", 16, isPowerOfFour(16))
	fmt.Printf("%d is %spower of four\n", 32, isPowerOfFour(32))
	fmt.Printf("%d is %spower of four\n", -65536, isPowerOfFour(-65536))
}