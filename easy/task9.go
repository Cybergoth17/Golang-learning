package main

import (
  "fmt"
  "strings"
)

func defangIPaddr(address string) string {
  res := strings.Replace(address, ".", "[.]", 4)
  return res
}
func main() {
  fmt.Println("Result: ")
  address := "1.1.1.1"
  fmt.Println(defangIPaddr(address))
}
