package  main

import (
    "fmt"
)

func printer(w []string) {
  for _, word := range w {
    fmt.Printf("%s,", word)
  }
  fmt.Printf("\n")
  w[2] = "blue"
}

func main() {
    words := []string{"the", "quick", "brown", "fox"}
    printer(words[1:])
    printer(words)

    word1 := make([]string, 4)
    word1[0] = "hello"
    word1[1] = "world"
    word1[2] = "good"
    word1[3] = "morning"
    printer(word1)
    word1 = append(word1, "go")
    printer(word1)

    word2 := make([]string, 4)
    copy(word2, word1)
    printer(word2)
    word2[1] = "你好"
    printer(word1)
    printer(word2)
}
