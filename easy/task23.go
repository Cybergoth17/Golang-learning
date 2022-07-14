package main

import (
  "fmt"
)

func balancedStringSplit(s string) int {
  count := 0
  bal := 0
  for i := 0; i < len(s); i++ {
    letter := string(s[i])
    if letter == "L" {
      bal++
    } else {
      bal--
    }
    if bal == 0 {
      count++
    }
  }
  return count
}
func main() {
  fmt.Println(balancedStringSplit("RLLLLRRRLR"))
}
