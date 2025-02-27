package main

import "fmt"

func fibonacci(x int) int {
	fmt.Println("Calculating F(", x, ")")
	defer fmt.Println("Done!")
	if x < 0 {
		return -1
	}
	if x == 0 {
		return 0
	}
	f0, f1 := 0, 1
	for i := 2; i <= x; i++ {
		f0, f1 = f1, f0+f1
	}
	return f1
}
