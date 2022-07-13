package main

import "fmt"

const ArraySize = 7

type HashTable struct {
  array [ArraySize]*bucket
}
type bucket struct {
  head *bucketNode
}
type bucketNode struct {
  key  string
  next *bucketNode
}

func (h *HashTable) Insert(key string) {
  index := hash(key)
  h.array[index].insert(key)
}
func (h *HashTable) Search(key string) bool {
  //index := hashkey)
  return false
}
func (h *HashTable) delete(key string) {
  //index := hash(key)
}
func hash(key string) int {
  sum := 0
  for _, v := range key {
    sum += int(v)
  }
  return sum % ArraySize
}
func Init() *HashTable {
  result := &HashTable{}
  for i := range result.array {
    result.array[i] = &bucket{}
  }
  return result
}
func (b *bucket) insert(k string) {
  newNode := &bucketNode{key: k}
  newNode.next = b.head
  b.head = newNode
}
func main() {
  hassh := Init()
  fmt.Println(hassh)
  fmt.Println(hash("RANDY"))
  hassh.Insert("Randy")
  fmt.Println(hassh)
}
