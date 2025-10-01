package main

import (
	"fmt"
	"sort"
	"strings"
)

func areAnagrams(s1, s2 string) bool {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	if len(s1) != len(s2) {
		return false
	}

	r1 := []rune(s1)
	r2 := []rune(s2)

	sort.Slice(r1, func(i, j int) bool { return r1[i] < r1[j] })
	sort.Slice(r2, func(i, j int) bool { return r2[i] < r2[j] })

	for i := range r1 {
		if r1[i] != r2[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(areAnagrams("listen", "silent"))
	fmt.Println(areAnagrams("hello", "world"))
	fmt.Println(areAnagrams("race", "care"))
}
