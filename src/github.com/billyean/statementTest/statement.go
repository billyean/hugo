package main

import (
	"fmt"
	"os"
)

func main() {
	numByte, err := fmt.Printf("Hello, world!")

	// If statement
	//if err != nil {
	//	os.Exit(1);
	//}
	//
	//fmt.Printf("read %d bytes\n", numByte)

	// Switch statement.
	switch {
	case err != nil:
		os.Exit(1)
	case numByte == 0:
		fmt.Printf("No bytes output\n")
	case numByte != 13:
		fmt.Printf("Wrong number of characters: %d\n", numByte)
	default:
		fmt.Printf("\nOK!\n")
	}
	fmt.Printf("End\n");

	// Switch variable staement.
	atoz := "I am an adroable boy!"

	var vowels int
	var otherC int
	for _, v := range atoz {
		switch v {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowels = vowels + 1
		default:
			otherC = otherC + 1
		}
	}

	fmt.Printf("Number of vowels characters are %d, other characters are %d", vowels, otherC)
}
