package main

import (
  "fmt"
  "math/rand"
)

func main() {
  arr := []int{8, 9, 71, 6, 85, 4, 2, 7}
  arr1 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
  fmt.Println("bubble sort")
  fmt.Println(bubbleSort(arr))
  fmt.Println("quick sort")
  fmt.Println(quickSort(arr1))
  fmt.Println("merge sort")
  fmt.Println(mergeSort(arr))
}
func bubbleSort(arr []int) []int {
  for i := 0; i < len(arr)-1; i++ {
    for j := 0; j < len(arr)-i-1; j++ {
      if arr[j] > arr[j+1] {
        arr[j], arr[j+1] = arr[j+1], arr[j]
      }
    }
  }
  return arr
}
func quickSort(arr []int) []int {
  if len(arr) < 2 {
    return arr
  }
  left, right := 0, len(arr)-1
  pivot := rand.Int() % len(arr)
  arr[pivot], arr[right] = arr[right], arr[pivot]
  for i, _ := range arr {
    if arr[i] < arr[right] {
      arr[left], arr[i] = arr[i], arr[left]
      left++
    }
  }
  arr[left], arr[right] = arr[right], arr[left]
  quickSort(arr[:left])
  quickSort(arr[left+1:])
  return arr
}
func mergeSort(arr []int) []int {
  num := len(arr)
  if num == 1 {
    return arr
  }
  middle := int(num / 2)
  left := make([]int, middle)
  right := make([]int, num-middle)
  for i := 0; i < num; i++ {
    if i < middle {
      left[i] = arr[i]
    } else {
      right[i-middle] = arr[i]
    }
  }
  return merge(mergeSort(left), mergeSort(right))
}
func merge(left, right []int) (result []int) {
  result = make([]int, len(left)+len(right))
  i := 0
  for len(left) > 0 && len(right) > 0 {
    if left[0] < right[0] {
      result[i] = left[0]
      left = left[1:]
    } else {
      result[i] = right[0]
      right = right[1:]
    }
    i++
  }
  for j := 0; j < len(left); j++ {
    result[i] = left[j]
    i++
  }
  for j := 0; j < len(right); j++ {
    result[i] = right[j]
    i++
  }
  return
}
