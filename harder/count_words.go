package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a sentence: ")
	sentence, _ := reader.ReadString('\n')

	// Remove leading/trailing spaces
	sentence = strings.TrimSpace(sentence)

	// Split by spaces
	words := strings.Fields(sentence)

	fmt.Println("Number of words:", len(words))
}
