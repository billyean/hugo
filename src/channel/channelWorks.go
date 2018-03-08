package channel

import (
	"math/rand"
)

func Emit(c chan string) {
	words := []string{"I", "am", "adorable"}

	for _, word := range words {
		c <- word
	}
	close(c)
}

func RandInts(c chan int) {
	count := 0
	for count < 100 {
		c <- rand.Intn(1 << 16)
		count += 1
	}
	close(c)
}