package main

import "fmt"

func main() {
	var word string
	fmt.Print("enter your string: ")
	fmt.Scanln(&word)
	reveresed := ""

	for i := 1; i <= len(word); i++ {
		last_letter := word[len(word)-i]

		reveresed = reveresed + string(last_letter)

	}
	fmt.Println(reveresed)
}
