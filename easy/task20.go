package main

import (
  "fmt"
  "strings"
)

func interpret(command string) string {
  command = strings.Replace(command, "()", "o", len(command))
  command = strings.Replace(command, "(", "", len(command))
  command = strings.Replace(command, ")", "", len(command))
  return command
}
func main() {
  fmt.Println(interpret("(al)G(al)()()G"))
}
