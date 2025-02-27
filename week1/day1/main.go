package main

import (
	"fmt"
)

func fib(n int) int {
	fmt.Printf("Finding the Fibonacci number at the position %d ...\n", n)
	defer fmt.Println("Done!")
	if n <= 1 {
		return n
	}
	n2, n1 := 0, 1
	for i := 2; i <= n; i++ {
		n2, n1 = n1, n1+n2
	}
	fmt.Printf("The Fibonacci number at position %d is %d\n", n, n1)
	return n1
}

func main() {
	n := 10 // Example value
	fib(n)
}
