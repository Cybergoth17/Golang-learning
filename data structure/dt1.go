package main

import (
  "fmt"
)

type node struct {
  data int
  next *node
}
type LinkedList struct {
  head   *node
  length int
}

func (l *LinkedList) prepend(n *node) {
  second := l.head
  l.head = n
  l.head.next = second
  l.length++
}
func (l LinkedList) printListData() {
  toPrint := l.head
  for l.length != 0 {
    fmt.Printf(" %d", toPrint.data)
    toPrint = toPrint.next
    l.length--
  }
}

func (l *LinkedList) delete(value int) {
  previous := l.head
  if l.length == 0 {
    return
  }
  if l.head.data == value {
    l.head = l.head.next
    l.length--
    return
  }
  for previous.next.data != value {
    if previous.next.next == nil {
      return
    }
    previous = previous.next
  }
  previous.next = previous.next.next
  l.length--
}
func main() {
  myList := &LinkedList{}
  node1 := &node{data: 48}
  node2 := &node{data: 78}
  node3 := &node{data: 58}
  node4 := &node{data: 10}
  myList.prepend(node1)
  myList.prepend(node2)
  myList.prepend(node3)
  myList.prepend(node4)
  fmt.Println("befor deleting")
  myList.printListData()
  myList.delete(48)
  fmt.Println(" after deleting")
  myList.printListData()
}
