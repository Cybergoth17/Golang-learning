package main

import (
  "fmt"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
  max := 0
  arr := make([]bool, len(candies))
  for i := 0; i < len(candies); i++ {
    max = maxx(max, candies[i])
    candies[i] += extraCandies
  }
  for i := 0; i < len(candies); i++ {
    if candies[i] >= max {
      arr[i] = true
    } else {
      arr[i] = false
    }
  }
  fmt.Println(candies)
  fmt.Println(arr)
  fmt.Println(max)
  return arr
}
func maxx(a int, b int) int {
  if a > b {
    return a
  }
  return b
}
func main() {
  fmt.Println("Result: ")
  arr := []int{2, 3, 5, 1, 3}
  fmt.Println(kidsWithCandies(arr, 3))
}
