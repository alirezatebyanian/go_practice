package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Print("enter your input: ")
	fmt.Scanln(&input)
	for j := 1; j <= 10; j++ {
		fmt.Print(input*j, " ")
	}
	fmt.Println()
}
