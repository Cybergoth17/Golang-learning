package main

import "fmt"

type animal interface {
  walker
  runner
}
type walker interface {
  walk()
}
type runner interface {
  run()
}
type bird interface {
  walker
  flier
}
type flier interface {
  fly()
}
type eagle struct{}
type cat struct{}

func (c *cat) walk() {
  fmt.Println(" cat is walking")
}
func (c *cat) run() {
  fmt.Println(" cat is running")
}
func (e *eagle) walk() {
  fmt.Println(" eagle is walking")
}
func (e *eagle) fly() {
  fmt.Println(" eagle is flying")
}
func walk(w walker) {
  w.walk()
}
func main() {
  var c animal = &cat{}
  var e bird = &eagle{}
  walk(e)
  walk(c)
  c.run()
}
