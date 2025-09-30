package main

import (
	"fmt"
	"slices"
)

// helper function: check if element exists in slice
func contains(slice []string, element string) int {
	for i, value := range slice {
		if value == element {
			return i
		}
	}
	return -1
}

// play one round of the game with a given word
func playWord(word []string, maxGuesses int) bool {
	letters := make([]string, len(word))
	copy(letters, word) // copy to avoid modifying original

	var guess string
	for i := 0; i < len(word)+maxGuesses; i++ {
		fmt.Println("Enter your guess:")
		fmt.Scanln(&guess)

		itemIndex := contains(letters, guess)
		if itemIndex >= 0 {
			fmt.Println("✅ Correct!")
			letters = slices.Delete(letters, itemIndex, itemIndex+1)
		} else {
			fmt.Println("❌ Wrong!")
		}

		// if all letters guessed, word is done
		if len(letters) == 0 {
			fmt.Println("🎉 You guessed the word!")
			return true
		}
	}

	fmt.Println("😢 You lost this word.")
	return false
}

func main() {
	// 3 different words (arrays of letters)
	words := [][]string{
		{"b", "a", "n", "a", "n", "a"},
		{"o", "r", "a", "n", "g", "e"},
		{"g", "r", "a", "p", "e"},
	}

	// play each word in order
	for round, word := range words {
		fmt.Printf("\n--- Round %d ---\n", round+1)
		if !playWord(word, 8) {
			fmt.Println("Game over!")
			return
		}
	}

	fmt.Println("🏆 Congratulations! You won all 3 rounds!")
}
