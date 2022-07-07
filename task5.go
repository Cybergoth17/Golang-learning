package main

import (
	"fmt"
)

func buildArray(nums []int) []int {
	result := make([]int, len(nums))

	for index, num := range nums {
		result[index] = nums[num]
	}

	return result
}
func main() {
	fmt.Println("Result: ")
}
