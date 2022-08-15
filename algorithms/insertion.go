package main

import (
  "fmt"
)

func main() {
  arr1 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
  fmt.Println("insertion sort")
  insertionSort(arr1)
  fmt.Println(arr1)
}
func insertionSort(arr []int) {
  n := len(arr)
  for i := 1; i < n; i++ {
    for j := i; j > 0; j-- {
      if arr[j-1] > arr[j] {
        arr[j-1], arr[j] = arr[j], arr[j-1]
      }
    }
  }
}
