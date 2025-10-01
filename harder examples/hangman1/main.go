package main

import (
	"fmt"
	"slices"
)

func contains(slice []string, element string) int {

	for i, value := range slice {
		if value == element {
			return i
		}
	}
	return -1 // -1 to mix up with indexes
}

func main() {
	letters := []string{"b", "a", "n", "a", "n", "a"}

	var guess string
	for i := 0; i < len(letters)+8; i++ {
		fmt.Println("enter your guess:")
		fmt.Scanln(&guess)

		item_index := contains(letters, guess)

		if item_index >= 0 {
			fmt.Println("True")
			letters = slices.Delete(letters, item_index, item_index+1)
		} else {
			fmt.Println("false")
		}

	}

	// final check
	if len(letters) == 0 {
		fmt.Println("you won")

	} else {
		fmt.Println("you lost")
	}
}
