package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var guess int
	fmt.Println("guess the number between 1 to 20")
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(20) + 1
	for i := 0; i < 7; i++ {

		fmt.Print("enter your guess: ")
		fmt.Scanln(&guess)

		if guess > r {
			fmt.Println("lower")
		} else if guess < r {
			fmt.Println("higher")
		} else {
			fmt.Println("you'r goddam right")
			break
		}
	}

}
