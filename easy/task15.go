package main

import (
  "fmt"
  "strings"
)

func numJewelsInStones(jewels string, stones string) int {
  count := 0
  if strings.ContainsAny(jewels, stones) {
    chars := []rune(stones)
    for i := 0; i < len(chars); i++ {
      char := string(chars[i])
      if strings.ContainsAny(jewels, char) {
        count++
      }
    }
  }
  return count
}
func main() {
  accounts := "z"
  cc := "ZZ"
  fmt.Println(numJewelsInStones(accounts, cc))
}
