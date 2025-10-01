package main

import (
	"fmt"
	"strings"
)

func vowelcase(s string) int {
	s = strings.ToLower(s)
	count := 0
	for _, ch := range s {
		if ch == 'a' || ch == 'o' || ch == 'e' || ch == 'u' || ch == 'i' {
			count++
		}
	}
	return count
}
func main() {
	var input string
	fmt.Print("enter your string: ")
	fmt.Scanln(&input)

	fmt.Println("the number of vowel that this string has is: ", vowelcase(input))
}
