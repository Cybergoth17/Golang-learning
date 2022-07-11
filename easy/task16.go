package main

import (
  "fmt"
)

func shuffle(nums []int, n int) []int {
  res := make([]int, 0)
  for i := 0; i < n; i++ {
    res = append(res, nums[i], nums[i+n])
  }
  return res
}
func main() {
  accounts := []int{1, 2, 3, 4, 5, 6}
  fmt.Println(shuffle(accounts, 3))
}
