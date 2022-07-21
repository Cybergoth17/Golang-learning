package main

import (
  "fmt"
)

func moveZeroes(nums []int) {
  j := 0
  for i, x := range nums {
    if x != 0 {
      tmp := nums[i]
      nums[i] = nums[j]
      nums[j] = tmp
      j++
    }
    i++
  }
}
func main() {
  arr := []int{7, 0, 1, 8, 0, 7}
  moveZeroes(arr)
  fmt.Println(arr)
}
