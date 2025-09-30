package main

import (
	"fmt"
)

func commonElements(a, b []int) []int {
	set := make(map[int]bool)
	result := []int{}

	// put all elements of 'a' into the map
	for _, v := range a {
		set[v] = true
	}

	// check if elements of 'b' exist in the map
	for _, v := range b {
		if set[v] {
			result = append(result, v)
			delete(set, v) // avoid duplicates in result
		}
	}

	return result
}

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{3, 4, 5, 6}

	common := commonElements(slice1, slice2)
	fmt.Println("Common elements:", common) // [3 4]
}
