package leetcode0748

import (
	"math"
	"strings"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	licensePlate = strings.ToLower(licensePlate)
	chars := [26]int{}
	for i := 0; i < len(licensePlate); i++ {
		b := licensePlate[i]
		if 'a' <= b && b <= 'z' {
			chars[b-'a']++
		}
	}

	match := func(word string) bool {
		cs := [26]int{}
		for i := 0; i < len(word); i++ {
			b := word[i]
			if 'a' <= b && b <= 'z' {
				cs[b-'a']++
			}
		}
		for i, c := range cs {
			if c < chars[i] {
				return false
			}
		}
		return true
	}

	min := math.MaxInt32
	res := ""
	for _, word := range words {
		if len(word) > min {
			continue
		}

		word = strings.ToLower(word)
		if match(word) && len(word) < min {
			min = len(word)
			res = word
		}
	}
	return res
}
