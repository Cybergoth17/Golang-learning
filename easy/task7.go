package main

import (
  "fmt"
  "strings"
)

func mostWordsFound(sentences []string) int {
  x := 0
  for i := 0; i < len(sentences); i++ {
    wordCount := strings.Count(sentences[i], " ")
    x = max(x, wordCount+1)
  }
  return x
}
func max(a int, b int) int {
  if a > b {
    return a
  }
  return b
}
func main() {
  fmt.Println("Result: ")
  accounts := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
  fmt.Println(mostWordsFound(accounts))
}
