package main

import (
	"fmt"
)

func main() {
	var n1 int
	var n2 int
	var s string
	for i := 0; i >= 0; i++ {
		fmt.Print("enter your first number: ")
		fmt.Scanln(&n1)
		fmt.Print("enter your second number: ")
		fmt.Scanln(&n2)
		fmt.Print("enter your sign that you want: ")
		fmt.Scanln(&s)
		if s == "+" {
			fmt.Println(n1 + n2)
		} else if s == "-" {
			fmt.Println(n1 - n2)
		} else if s == "*" {
			fmt.Println(n1 * n2)
		} else if s == "/" {
			fmt.Println(n1 / n2)
		} else {
			fmt.Println("the sign is invalid!")
		}
	}

}
