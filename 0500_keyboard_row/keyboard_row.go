package leetcode0500

import (
	"regexp"
)

func findWords(words []string) []string {
	r := regexp.MustCompile("(?i)^[qwertyuiop]*$|^[asdfghjkl]*$|^[zxcvbnm]*$")
	res := []string{}

	for _, word := range words {
		if r.MatchString(word) {
			res = append(res, word)
		}
	}
	return res
}
