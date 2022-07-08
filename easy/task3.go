package main

import (
	"fmt"
	"sort"
)

func sortArray(nums []int) []int {
	sort.Sort(sort.IntSlice(nums))
	return nums
}
func main() {
	fmt.Println("Result: ")
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(sortArray(nums))
}
