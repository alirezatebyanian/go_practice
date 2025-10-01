package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	resualt := 1
	for i := 2; i <= n; i++ {
		resualt = resualt * i
	}
	return resualt
}

func main() {
	var input int
	fmt.Print("enter the fac you want to calculate: ")
	fmt.Scanln(&input)
	fmt.Println("factorial of ", input, " is ", factorial(input))

}
