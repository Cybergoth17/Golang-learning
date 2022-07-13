package main

import (
  "fmt"
  "strconv"
  "strings"
)

const space = " "

func sortSentence(s string) string {
  kdd := strings.Split(s, space)
  khh := make([]string, len(kdd))
  for _, word := range kdd {
    m, _ := strconv.Atoi(word[len(word)-1:])
    khh[m-1] = word[:len(word)-1]
  }
  return strings.Join(khh, space)
}

func main() {
  s := "Myself2 Me1 I4 and3"
  fmt.Println(sortSentence(s))
}
