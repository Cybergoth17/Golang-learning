package main

import (
  "fmt"
)

func finalValueAfterOperations(operations []string) int {
  x := 0
  for i := 0; i < len(operations); i++ {
    if operations[i] == "X++" {
      x++
    } else if operations[i] == "++X" {
      x++
    } else {
      x--
    }
  }
  return x
}
func main() {
  fmt.Println("Result: ")
  operations := []string{"++X", "++X", "X++"}
  fmt.Println(finalValueAfterOperations(operations))
}
