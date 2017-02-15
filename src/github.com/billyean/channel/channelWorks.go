package main

import (
	"fmt"
	"math/rand"
)

func emit(c chan string) {
	words := []string{"I", "am", "adorable"}

	for _, word := range words {
		c <- word
	}
	close(c)
}

func randInts(c chan int) {
	count := 0
	for count < 100 {
		c <- rand.Intn(1 << 16)
		count += 1
	}
	close(c)
}

func main() {
	wordsChannel := make(chan string)
	intsChannel := make(chan int)
	go randInts(intsChannel)
	go emit(wordsChannel)
	for word := range wordsChannel {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")
	for num := range intsChannel {
		fmt.Printf("%d\n", num)
	}
}
