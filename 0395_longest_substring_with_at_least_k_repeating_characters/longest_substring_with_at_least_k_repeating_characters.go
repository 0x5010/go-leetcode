package leetcode0395

func longestSubstring(s string, k int) int {
	n := len(s)
	m := make([]int, 26)

	res := 0
	for h := 1; h <= 26; h++ {
		for i := range m {
			m[i] = 0
		}
		i, j := 0, 0
		unique, noLessThanK := 0, 0
		for j < n {
			if unique <= h {
				index := s[j] - 'a'
				if m[index] == 0 {
					unique++
				}
				m[index]++
				if m[index] == k {
					noLessThanK++
				}
				j++
			} else {
				index := s[i] - 'a'
				if m[index] == k {
					noLessThanK--
				}
				m[index]--
				if m[index] == 0 {
					unique--
				}
				i++
			}

			if unique == h && unique == noLessThanK && j-i > res {
				res = j - i

			}
		}
	}
	return res
}
