package leetcode0336

func palindromePairs(words []string) [][]int {
	n := len(words)
	if n < 2 {
		return nil
	}
	res := [][]int{}

	m := make(map[string]int, n)
	for i := 0; i < n; i++ {
		m[words[i]] = i
	}

	for i, word := range words {
		for j := 0; j <= len(word); j++ {
			l, r := word[:j], word[j:]

			if isPalindrome(l) {
				if k, ok := m[reverse(r)]; ok && i != k {
					res = append(res, []int{k, i})
				}
			}

			if len(r) != 0 && isPalindrome(r) {
				if k, ok := m[reverse(l)]; ok && i != k {
					res = append(res, []int{i, k})
				}
			}
		}
	}
	return res
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func reverse(s string) string {
	bs, i, j := []byte(s), 0, len(s)-1
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return string(bs)
}
