package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i <= len(nums)-1; i++ {
		for z := i + 1; z <= len(nums)-1; z++ {
			result := nums[i] + nums[z]
			if result == target {
				g := []int{i, z}
				return g
			}
		}

	}
	return nil
}
func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
