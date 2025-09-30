package main

import (
	"fmt"
)

func main() {
	var num1 int
	var num2 int
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	if num1 > num2 {
		fmt.Println("the larger one is ", num1)
	} else if num2 == num1 {
		fmt.Println("none of them are larger")
	} else {
		fmt.Println("the larger one is ", num2)
	}

}
