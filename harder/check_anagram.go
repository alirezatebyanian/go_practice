package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var input1 string
	var input2 string
	fmt.Print("enter the first string: ")
	fmt.Scanln(&input1)
	fmt.Print("enter the second string: ")
	fmt.Scanln(&input2)
	fmt.Println(areAnagrams(input1, input2))
}

func areAnagrams(s1, s2 string) bool {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	if len(s1) != len(s2) {
		return false
	}

	runes1 := []rune(s1)
	runes2 := []rune(s2)

	sort.Slice(runes1, func(i, j int) bool { return runes1[i] < runes1[j] })
	sort.Slice(runes2, func(i, j int) bool { return runes2[i] < runes2[j] })

	for i := range runes1 {
		if runes1[i] != runes2[i] {
			return false
		}
	}

	return true
}
