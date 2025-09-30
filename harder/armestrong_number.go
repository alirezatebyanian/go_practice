package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if isArmstrong(num) {
		fmt.Print(num, "is an Armstrong number.")
	} else {
		fmt.Print(num, "is not an Armstrong number.")
	}
}

func isArmstrong(number int) bool {
	original := number
	sum := 0

	// تعداد ارقام
	digits := 0
	temp := number
	for temp != 0 {
		digits++
		temp /= 10
	}

	// جمع ارقام به توان تعداد ارقام
	temp = number
	for temp != 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}

	return sum == original
}
