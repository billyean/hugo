package main

import (
	"fmt"
)

func emit(c chan string) {
	words := []string{"I", "am", "adorable"}

	for _, word := range words {
		c <- word
	}

	close(c)
}

func main() {
	wordsChannel := make(chan string)
	go emit(wordsChannel)
	for word := range wordsChannel {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")
}
