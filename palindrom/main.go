package main

import (
	"fmt"
)

func main() {
	var word string
	fmt.Print("enter your word: ")
	fmt.Scanln(&word)
	reveresed := ""

	for i := 1; i <= len(word); i++ {
		last_letter := word[len(word)-i]

		reveresed += string(last_letter)
	}

	if word == reveresed {
		fmt.Println("the word is Palidrom")
	} else {
		fmt.Println("the word is not Palidrom")
	}

}
