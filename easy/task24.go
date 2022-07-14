package main

import (
  "fmt"
  "sort"
)

func minimumSum(num int) int {
  nums := make([]int, 0, 0)
  for i := 0; i < 4; i++ {
    nums = append(nums, num%10)
    num /= 10
  }
  sort.Ints(nums)
  return nums[0]*10 + nums[1]*10 + nums[2] + nums[3]
}

func main() {
  num := 2358
  fmt.Println(minimumSum(num))
}
