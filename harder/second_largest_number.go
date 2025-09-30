package main

import (
	"fmt"
)

func secondLargest(nums []int) (int, bool) {
	if len(nums) < 2 {
		return 0, false
	}

	largest := nums[0]
	second := -1 << 31

	for _, n := range nums {
		if n > largest {
			second = largest
			largest = n
		} else if n > second && n < largest {
			second = n
		}
	}

	if second == -1<<31 {
		return 0, false
	}
	return second, true
}

func main() {
	slice := []int{10, 20, 4, 45, 99, 99, 32}
	if sec, ok := secondLargest(slice); ok {
		fmt.Println("Second largest:", sec)
	} else {
		fmt.Println("No second largest number found.")
	}
}
