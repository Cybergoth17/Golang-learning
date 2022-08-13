package main

import (
  "fmt"
)

func main() {
  arr1 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
  fmt.Println("binary search")
  fmt.Println(binarySearch(5, arr1))
  fmt.Println(binarySearchBool(1, arr1))
}

func binarySearch(value int, a []int) int {
  low, high := 0, len(a)-1
  for low <= high {
    mid := (high + low) / 2
    if value == a[mid] {
      return mid
    } else if value < a[mid] {
      high = mid - 1
    } else {
      low = mid + 1
    }
  }
  return -1
}
func binarySearchBool(needle int, haystack []int) bool {

  low := 0
  high := len(haystack) - 1

  for low <= high {
    median := (low + high) / 2

    if haystack[median] < needle {
      low = median + 1
    } else {
      high = median - 1
    }
  }

  if low == len(haystack) || haystack[low] != needle {
    return false
  }

  return true
}
