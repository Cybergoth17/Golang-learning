package main

import (
  "fmt"
)

func runningSum(nums []int) []int {
  arr := make([]int, len(nums))
  elem := 0
  for i := 0; i <= len(nums)-1; i++ {
    elem += nums[i]
    arr[i] = elem
  }
  return arr
}
func main() {
  fmt.Println("Result: ")
  nums := []int{1, 2, 3, 4}
  fmt.Println(runningSum(nums))
}
