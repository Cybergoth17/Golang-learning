package main

import (
	"fmt"
)

func getConcatenation(nums []int) []int {
	arr := make([]int, len(nums)<<1)
	copy(arr, nums)
	copy(arr[len(nums):], nums)
	return arr
}
func main() {
	fmt.Println("Result: ")
	nums := []int{1, 5, 8}
	fmt.Println(getConcatenation(nums))
}
