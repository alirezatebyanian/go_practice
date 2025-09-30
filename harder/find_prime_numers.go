package main

import (
	"fmt"
)

func main() {
	for num := 2; num <= 100; num++ {
		if isPrime(num) {
			fmt.Print(num, " ")
		}
	}
}

// Function to check if a number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
