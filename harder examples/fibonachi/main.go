package main

import (
	"fmt"
)

func main() {

	fib_nums := []int{0, 1}
	var n int
	fmt.Print("enter the end: ")
	fmt.Scanln(&n)
	for i := 1; i <= n-2; i++ {
		num := fib_nums[i]
		privious_num := fib_nums[i-1]
		fib_nums = append(fib_nums, num+privious_num)
	}

	fmt.Println(fib_nums)

}
