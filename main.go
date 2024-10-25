package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Print("factorial(9) = ")
	fmt.Println(factorial(9))

	fmt.Print("fibonacci(9) = ")
	fmt.Println(fibonacci(9))
}
