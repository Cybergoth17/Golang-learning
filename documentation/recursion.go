package main

import "fmt"

func main() {
   fmt.Println(factorial(5))
   fmt.Println(fib(7))
}
func factorial(n int) int{
    if n<=1{
        return 1
    }
    return n*factorial(n-1)
}
func fib(n int) int{
    if n<2{
        return n
    }
    return fib(n-1)+fib(n-2)
}
