package leetcode0030

func findSubstring(s string, words []string) []int {
	sLen, wordsLen := len(s), len(words)
	if sLen == 0 || wordsLen == 0 {
		return nil
	}
	wLen := len(words[0])
	if sLen < wordsLen*wLen {
		return nil
	}

	count := map[string]int{}
	for _, w := range words {
		count[w]++
	}
	indexes := []int{}
	for i := 0; i < sLen-wordsLen*wLen+1; i++ {
		seen := map[string]int{}
		j := 0
		for j < wordsLen {
			word := s[i+j*wLen : i+(j+1)*wLen]
			if _, ok := count[word]; ok {
				seen[word]++
				if seen[word] > count[word] {
					break
				}
			} else {
				break
			}
			j++
		}
		if j == wordsLen {
			indexes = append(indexes, i)
		}
	}
	return indexes
}
