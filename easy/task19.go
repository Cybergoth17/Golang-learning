package main

import "fmt"

func decode(encoded []int, first int) []int {
  decode := make([]int, 0, len(encoded)+1)
  decode = append(decode, first)
  for i := range encoded {
    decode = append(decode, decode[i]^encoded[i])
  }
  return decode
}
func main() {
  arr := []int{6, 2, 7, 3}
  fmt.Println(decode(arr, 4))
}
