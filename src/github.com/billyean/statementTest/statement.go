package main

import (
	"fmt"
	"os"
)

func getPrefix(name string) (prefix string) {
	switch name {
	case "Tristan":
		prefix = "Mr"
	case "Tina":
		prefix = "Mrs"
	case "Ruby":
		prefix = "Miss"
	case "Zachary", "Rachael":
		prefix = "Buddy"
	}
	return
}

func callNTime(times int, message string) {
	for i := 0; i < times; i++ {
		fmt.Println(message)
	}
}

func callNTimeWhile(times int, message string) {
	var i = 0
	for i < times{
		fmt.Println(message)
		i++
	}
}

func callNTimeInfinite1(times int, message string) {
	var i = 0
	for {
		fmt.Println(message)
		i++
		if i > times {
			break
		}
	}
}

func callNTimeInfinite2(times int, message string) {
	var i = 0
	for {
		fmt.Println(message)
		i++
		if i <= times {
			continue
		}
	}
}



func callNTimeRangee(times int, message string) {

}

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

	//yan := {"Tristan", "Tina", "Ruby", "Zacahry"}
	//for name = range yan {
	//	fmt.Println()
	//}

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
