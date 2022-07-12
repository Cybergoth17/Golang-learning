package main

import (
  "fmt"
  "sort"
)

func main() {
  array := []float64{5.58, 7.89, 1.12}
  arr := []int{4, 5, 6, 7, 0, 2, 1, 3}
  // int are sorted return false if arr not sorted
  fmt.Println(sort.IntsAreSorted(arr))
  // Ints sort array in increasing order
  sort.Ints(arr)
  // Float64s sort float array in increasing order
  sort.Float64s(array)
  //check array are sorted or not, return boolean value
  fmt.Println(sort.Float64sAreSorted(array))
  //search method search your input target
  i := sort.Search(len(arr), func(i int) bool { return arr[i] >= 7 })
  if i < len(arr) && arr[i] == 7 {
    fmt.Printf("found 7 at index %d in %v\n", i, arr)
  } else {
    fmt.Printf("not found")
  }
  //search floats64s will find target element if we input
  x := 7.89
  y := sort.SearchFloat64s(array, x)
  fmt.Printf("element %g found at index %d\n", x, y)
  //search integer value and return its own index
  e := sort.SearchInts(arr, 4)
  fmt.Printf("element 4 found at index %d\n", e)
  //strings method sort string slice by increasing order
  arrr := []string{"za", "e", "bal"}
  sort.Strings(arrr)
  fmt.Println(arrr)
  fmt.Println(array)
  fmt.Println(arr)
  sort.Sort(sort.Reverse(sort.IntSlice(arr)))
  fmt.Println(arr)
}
