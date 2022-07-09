package main

import (
  "fmt"
)

func numberOfSteps(num int) int {
  count := 0
  for num > 0 {
    if num%2 == 0 {
      num /= 2
      count++
    } else {
      num -= 1
      count++
    }
  }
  return count
}
func main() {
  fmt.Println("Result: ")
  nums := 123
  fmt.Println(numberOfSteps(nums))
}
