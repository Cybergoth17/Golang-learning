package main

import (
  "fmt"
  "strings"
)

func main() {
  //clone implement
  str := " I am nothing"
  a := strings.Clone(str)
  // contain implement
  b := strings.Contains(str, "am")
  // containsAny
  c := strings.ContainsAny(str, "um")
  //containsRune
  d := strings.ContainsRune(str, 97)
  //count
  e := strings.Count(str, "i")
  // cut
  f, g, h := strings.Cut(str, "ing")
  // equalFold
  i := strings.EqualFold(a, "i am nothing")
  // splits
  j := strings.Split(str, " ")
  // map
  rot13 := func(r rune) rune {
    switch {
    case r >= 'A' && r <= 'Z':
      return 'A' + (r-'A'+13)%26
    case r >= 'a' && r <= 'z':
      return 'a' + (r-'a'+13)%26
    }
    return r
  }
  fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
  fmt.Println(e)
  fmt.Println(f, g, h)
  fmt.Println(i)
  fmt.Println(j)
}
