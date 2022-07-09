package main

import (
  "fmt"
)

func subtractProductAndSum(n int) int {
  max := 1
  min := 0
  for 0 < n {
    digit := n % 10
    n = (n - digit) / 10
    max *= digit
    min += digit
  }
  return max - min
}
func main() {
  fmt.Println("Result: ")
  n := 114
  fmt.Println(subtractProductAndSum(n))
}
