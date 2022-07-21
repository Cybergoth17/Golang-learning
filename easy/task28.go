package main

import (
  "fmt"
  "sort"
)

func singleNumber(nums []int) int {
  sort.Ints(nums)
  return nums[len(nums)/2]
}
func main() {
  arr := []int{3, 3, 2}
  fmt.Println(singleNumber(arr))
}
