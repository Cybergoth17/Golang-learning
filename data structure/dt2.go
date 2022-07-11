package main

import "fmt"

type MaxHeap struct {
  array []int
}

func (h *MaxHeap) insert(key int) {
  h.array = append(h.array, key)
  h.maxHeapify(len(h.array) - 1)
}
func (h *MaxHeap) maxHeapify(index int) {
  for h.array[parent(index)] < h.array[index] {
    h.swap(parent(index), index)
    index = parent(index)
  }
}
func (h *MaxHeap) Extract() int {
  extracted := h.array[0]
  if len(h.array)-1 == 0 {
    fmt.Println("cannot extracted")
    return -1
  }
  h.array[0] = h.array[len(h.array)-1]
  h.array = h.array[:len(h.array)-1]
  h.maxHeapifyDown(0)
  return extracted
}
func (h *MaxHeap) maxHeapifyDown(index int) {
  lastIndex := len(h.array) - 1
  l, r := left(index), right(index)
  childToCompare := 0
  for l <= lastIndex {
    if l == lastIndex {
      childToCompare = l
    } else if h.array[l] > h.array[r] {
      childToCompare = l
    } else {
      childToCompare = r
    }
    if h.array[index] < h.array[childToCompare] {
      h.swap(index, childToCompare)
      index = childToCompare
      l, r = left(index), right(index)
    } else {
      return
    }
  }
}
func parent(i int) int {
  return (i - 1) / 2
}
func left(i int) int {
  return (i * 2) + 1
}
func right(i int) int {
  return (i * 2) + 2
}
func (h *MaxHeap) swap(i1, i2 int) {
  h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
func main() {
  m := &MaxHeap{}
  fmt.Println(m)
  buildHeap := []int{10, 20, 30, 5, 7, 9, 8}
  for _, v := range buildHeap {
    m.insert(v)
    fmt.Println(m)
  }
  for i := 0; i < 5; i++ {
    m.Extract()
    fmt.Println(m)
  }
}
