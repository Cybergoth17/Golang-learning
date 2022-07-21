package main

import (
  "fmt"
)

func moveZeroes(nums []int) {
  for i := 0; i < len(nums); i++ {
    if nums[i] == 0 {
      copy(nums[i:], nums[i+1:])
      nums[len(nums)-1] = 0
      nums = nums[:len(nums)-1]
    }
  }
}
func main() {
  arr := []int{7, 0, 1, 8, 0, 7}
  moveZeroes(arr)
  fmt.Println(arr)
}
