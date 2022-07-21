package main

import (
  "fmt"
)

func singleNumber(nums []int) int {
  m := map[int]int{}
  for i := 0; i < len(nums); i++ {
    m[nums[i]]++
  }
  for key, value := range m {
    if value == 1 {
      return key
    }
  }
  return 0
}
func main() {
  arr := []int{2, 7, 7, 2, 8}
  fmt.Println(singleNumber(arr))
}
