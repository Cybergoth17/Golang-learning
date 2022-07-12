package main

import (
  "fmt"
  "strings"
)

func restoreString(s string, indices []int) string {
  pizdos := make([]string, len(indices))
  for index, indice := range indices {
    pizdos[indice] = string(s[index])
  }
  return strings.Join(pizdos, "")
}
func main() {
  arr := []int{4, 5, 6, 7, 0, 2, 1, 3}
  fmt.Println(restoreString("codeleet", arr))
}
