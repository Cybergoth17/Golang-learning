package main

import "fmt"

type Math interface {
  Area() int
}
type Rectangle struct {
  Width  int
  Height int
}
type Square struct {
  Side int
}

func (r *Rectangle) Area() int {
  return (r.Height + r.Width) * 2
}
func (s *Square) Area() int {
  return s.Side * s.Side
}
func main() {
  s := &Square{5}
  r := &Rectangle{10, 2}
  fmt.Println(s.Area())
  fmt.Println(r.Area())
}
