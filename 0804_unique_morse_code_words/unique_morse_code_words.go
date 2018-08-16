package leetcode0804

import "strings"

func uniqueMorseRepresentations(words []string) int {
	t := [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := map[string]struct{}{}
	for _, word := range words {
		var bs strings.Builder
		for i := 0; i < len(word); i++ {
			bs.WriteString(t[word[i]-'a'])
		}
		m[bs.String()] = struct{}{}
	}
	return len(m)
}
