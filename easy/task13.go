package main

import (
  "fmt"
)

func smallerNumbersThanCurrent(nums []int) []int {
  var count []int
  for _, v := range nums {
    k := 0
    for _, v1 := range nums {
      if v1 != v {
        if v1 < v {
          k++
        }
      }
    }
    count = append(count, k)
  }
  return count
}
func main() {
  fmt.Println("Result: ")
  arr := []int{8, 1, 2, 2, 3}
  fmt.Println(smallerNumbersThanCurrent(arr))
}
