package main

import (
	"fmt"
)

func sum(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum = sum + i
	}
	return sum
}
func main() {
	var n int
	fmt.Print("insert the number: ")
	fmt.Scanln(&n)
	fmt.Println("the sum of 1 to ", n, "is ", sum(n))

}
