package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 2, 3, 4, 4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(slice []int) []int {
	seen := make(map[int]bool)
	var result []int

	for _, num := range slice {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}

	return result
}
