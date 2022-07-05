package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if reverse_int(x) == x {
		return true
	}
	return false
}
func reverse_int(n int) int {
	new_int := 0
	for n > 0 {
		remainder := n % 10
		new_int *= 10
		new_int += remainder
		n /= 10
	}
	return new_int
}

func main() {
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(121))
}
