package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
  c := 0
  for len(sandwiches) != 0 {
    stud, sand := students[0], sandwiches[0]
    if stud != sand {
      students = students[1:]
      students = append(students, stud)
      c++
    } else {
      students, sandwiches = students[1:], sandwiches[1:]
      c = 0
    }
    if c == len(students) {
      return c
    }
  }
  return 0
}

func main() {
  queue := []int{1, 1, 0, 0}
  stack := []int{0, 1, 0, 1}
  fmt.Println(countStudents(queue, stack))
}
