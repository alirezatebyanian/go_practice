package main

import (
	"fmt"
)

func findMax(nums []int) int {

	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}

	}
	return max
}

func main() {
	numbers := []int{5, 2, 9, 1, 7, 3}
	fmt.Println("Slice:", numbers)
	fmt.Println("Largest number:", findMax(numbers))
}
