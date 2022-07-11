package main

import (
  "fmt"
)

func maximumWealth(accounts [][]int) int {
  max := 0
  for _, account := range accounts {
    max = comp(max, sum(account))
  }
  return max
}
func sum(num []int) int {
  res := 0
  for _, item := range num {
    res += item
  }
  return res
}
func comp(a int, b int) int {
  if a > b {
    return a
  } else {
    return b
  }
}
func main() {
  accounts := [][]int{{1, 2, 3}, {5, 6, 7}}
  fmt.Println(maximumWealth(accounts))
}
