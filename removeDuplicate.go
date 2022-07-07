package main

import (
  "fmt"
  "sort"
)

func removeElement(nums []int, val int) int {
  j := 0
  newArr := make([]int, 0)
  for i := 0; i < len(nums); i++ {
    newArr = append(newArr, nums[i])
  }
  nums = nil
  for i := 0; i <= len(newArr)-1; i++ {
    if val != newArr[i] {
      nums = append(nums, newArr[i])
      j++
    }
  }
  sort.Sort(sort.IntSlice(nums))
  fmt.Println(nums)
  return j
}
func main() {
  fmt.Println("Result: ")
  nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
  fmt.Println(removeElement(nums, 3))
}