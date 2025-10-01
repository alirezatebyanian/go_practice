package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var input int
	fmt.Print("enter a number: ")
	fmt.Scanln(&input)

	str_input := strconv.Itoa(input)
	tavan := len(str_input)

	final_sum := 0

	for _, digit := range str_input {
		// each_digit_power := math.Pow(float64(digit), float64(tavan))

		int_digit, err := strconv.Atoi(string(digit))
		if err != nil {
			// ... handle error
			panic(err)
		}

		each_digit_powered := math.Pow(float64(int_digit), float64(tavan))
		final_sum = final_sum + int(each_digit_powered)

	}
	if input == final_sum {
		fmt.Println("it is anagram")
	} else {
		fmt.Println("it is not anagram")
	}

}
