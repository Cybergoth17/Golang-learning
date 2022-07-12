package main

import (
  "fmt"
)

type Stack struct {
  items []int
}

func (s *Stack) Push(val int) {
  s.items = append(s.items, val)
}
func (s *Stack) Pop(val int) int {
  l := len(s.items) - 1
  removed := s.items[l]
  s.items = s.items[:l]
  return removed
}
func main() {
  myStack := &Stack{}
  myStack.Push(100)
  myStack.Push(50)
  myStack.Push(200)
  fmt.Println(myStack)
}
