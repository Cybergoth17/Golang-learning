package main

import "fmt"

const Size = 7

type Hashtable struct {
  array [Size]*bucket
}
type bucket struct {
  head *bucketNode
}
type bucketNode struct {
  key  string
  next *bucketNode
}

func Init() *Hashtable {
  result := &Hashtable{}
  for i := range result.array {
    result.array[i] = &bucket{}
  }
  return result
}
func hash(key string) int {
  sum := 0
  for _, v := range key {
    sum += int(v)
  }
  return sum % Size
}
func (h *Hashtable) Insert(key string) {
  index := hash(key)
  h.array[index].insert(key)
}
func (h *Hashtable) Delete(key string) {
  index := hash(key)
  h.array[index].delete(key)
}
func (h *Hashtable) Search(key string) bool {
  index := hash(key)
  return h.array[index].search(key)
}
func (b *bucket) insert(k string) {
  if !b.search(k) {
    newNode := &bucketNode{key: k}
    newNode.next = b.head
    b.head = newNode
  } else {
    fmt.Println(k, " already exist")
  }
}
func (b *bucket) delete(k string) {
  if b.head.key == k {
    b.head = b.head.next
    return
  }
  previousNode := b.head
  for previousNode != nil {
    if previousNode.next.key == k {
      previousNode.next = previousNode.next.next
    }
    previousNode = previousNode.next
  }
}
func (b *bucket) search(k string) bool {
  currentNode := b.head
  for currentNode != nil {
    if currentNode.key == k {
      return true
    }
    currentNode = currentNode.next
  }
  return false
}
func main() {
  testHashTable := Init()
  list := []string{
    "ERIC",
    "KENNY",
    "KYLE",
    "STAN",
  }
  for _, v := range list {
    testHashTable.Insert(v)
  }
  testHashTable.Delete("STAN")
  fmt.Println("STAN ", testHashTable.Search("STAN"))
  fmt.Println("KENNY ", testHashTable.Search("KENNY"))
}
