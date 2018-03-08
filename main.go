package main

import "fmt"
import "array"

func demoBool() {
	var message = "Hello Tristan...\t"
	a, b, c := true,20, true
	fmt.Print(message, a, b, c)

	var say *string = &message
	fmt.Print("\n", *say)
}

func main() {
	demoBool()
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
	bitsCount := array.Count(17)
	fmt.Printf("[")

	i := 0

	for i < len(bitsCount) {
		fmt.Printf(" %d ", bitsCount[i])
		i = i + 1
	}
	fmt.Printf("]\n")
}
