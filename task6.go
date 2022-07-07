package main

import (
	"fmt"
)

func sum(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}
func main() {
	fmt.Println("Result: ")
	fmt.Println(sum(5, -8))
}
