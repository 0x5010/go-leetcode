package leetcode0966

import "strings"

func spellchecker(wordlist []string, queries []string) []string {
	words := map[string]struct{}{}
	cap, vowel := map[string]string{}, map[string]string{}
	for _, word := range wordlist {
		words[word] = struct{}{}
		lower := strings.ToLower(word)
		if _, ok := cap[lower]; !ok {
			cap[lower] = word
			if _, ok = vowel[todev(lower)]; !ok {
				vowel[todev(lower)] = word
			}
		}
	}
	for i, query := range queries {
		if _, ok := words[query]; ok {
			continue
		}
		lower := strings.ToLower(query)
		if word, ok := cap[lower]; ok {
			queries[i] = word
		} else if word, ok = vowel[todev(lower)]; ok {
			queries[i] = word
		} else {
			queries[i] = ""
		}
	}
	return queries
}

func todev(s string) string {
	bs := []byte(s)
	for i, b := range bs {
		if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' {
			bs[i] = '#'
		}
	}
	return string(bs)
}
