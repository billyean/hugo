package function

import (
	"fmt"
)

type Salutation struct {
    name string
    greeting string
}

func Greet(salutation Salutation) {
	fmt.Printf("%s %s", salutation.greeting, salutation.name)
}

// I am extremely no happy that abs(int) is not a supported
// function in Go, I also noticed people are talking it's not
// worth to be added since it's too generic.
// This won't make GO a better language than C++ because we need
// re-invent the wheel all the time.
func Abs(v int) int {
	if v >= 0 {
		return v
	} else {
		return -v
	}
}

// LeetCode exercise : power of 4 is a trick function
// https://leetcode.com/problems/power-of-four/
func IsPowerOfFour(v int) bool{
	v1 := Abs(v)
	return v1 & (v1 - 1) == 0 && (v1 - 1) % 3 == 0
}

//func main() {
//	fmt.Printf("%d is %spower of four\n", 5, IsPowerOfFour(5))
//	fmt.Printf("%d is %spower of four\n", 16, isPowerOfFour(16))
//	fmt.Printf("%d is %spower of four\n", 32, isPowerOfFour(32))
//	fmt.Printf("%d is %spower of four\n", -65536, isPowerOfFour(-65536))
//
//	var s = Salutation{"Bob", "Hello"}
//	greet(s)
//}