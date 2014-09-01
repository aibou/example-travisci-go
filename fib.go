package main

import "fmt"

func main() {
	fmt.Println(fib(6))
}

func fib(i int) int {
	if i < 1 {
		return 0
	}
	
	if i == 1 {
		return 1
	}

	return fib(i-1) + fib(i-2)
}
