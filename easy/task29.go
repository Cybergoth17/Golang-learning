package main

import "fmt"

type ListNode struct {
  Val  int
  Next *ListNode
}
type LinkedList struct {
  Len  int
  Head *ListNode
}

func (l LinkedList) print() {
  toPrint := l.Head
  for l.Len != 0 {
    fmt.Printf("%d", toPrint.Val)
    toPrint = toPrint.Next
    l.Len--
    fmt.Println("")
  }
}
func (l *LinkedList) Reverse() *ListNode {
  if l.Head == nil {
    return nil
  }
  curr := l.Head
  for curr.Next != nil {
    nextNode := curr.Next
    curr.Next = nextNode.Next
    nextNode.Next = l.Head
    l.Head = nextNode
  }
  return l.Head
}
func (l *LinkedList) add(key *ListNode) {
  second := l.Head
  l.Head = key
  l.Head.Next = second
  l.Len++
}

func main() {
  first := &LinkedList{}
  node1 := &ListNode{Val: 8}
  node2 := &ListNode{Val: 5}
  node3 := &ListNode{Val: 3}
  node4 := &ListNode{Val: 156}
  first.add(node1)
  first.add(node2)
  first.add(node3)
  first.add(node4)
  first.print()
  first.Reverse()
  first.print()
}
