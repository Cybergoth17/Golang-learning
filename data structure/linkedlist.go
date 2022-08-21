import "fmt"

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
func (l *LinkedList) show() {
  printed := l.head
  for l.length != 0 {
    fmt.Printf("%d ", printed.data)
    printed = printed.next
    l.length--
  }
  fmt.Printf("\n")
}
func (l *LinkedList) delete(value int) {
  if l.length == 0 {
    return
  }
  if l.head.data == value {
    l.head = l.head.next
  }
  deleted := l.head
  for deleted.next.data != value {
    if deleted.next.next == nil {
      return
    }
    deleted = deleted.next
  }
  deleted.next = deleted.next.next
  l.length--
}

func main() {
  list := LinkedList{}
  node1 := &node{data: 48}
  node2 := &node{data: 16}
  node3 := &node{data: 18}
  list.prepend(node1)
  list.prepend(node2)
  list.prepend(node3)
  list.delete(16)
  empty := LinkedList{}
  empty.delete(45)
  list.show()
}
