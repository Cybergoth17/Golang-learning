package main

import (
  "fmt"
)

func numIdenticalPairs(nums []int) int {
  pair := 0
  for i := 0; i <= len(nums)-1; i++ {
    for j := 1; j <= len(nums)-1; j++ {
      if nums[i] == nums[j] && i < j {
        pair++
      }
    }
  }
  return pair
}
func main() {
  fmt.Println("Result: ")
  nums := []int{1, 1, 1, 1}
  fmt.Println(numIdenticalPairs(nums))
}
