package leetcode0032

func longestValidParentheses(s string) int {
	l, r, maxLen, sLen := 0, 0, 0, len(s)
	for i := 0; i < sLen; i++ {
		if s[i] == '(' {
			l++
		} else {
			r++
		}
		if l == r {
			if 2*r > maxLen {
				maxLen = 2 * r
			}
		} else if r > l {
			l, r = 0, 0

		}
	}
	l, r = 0, 0
	for i := sLen - 1; i >= 0; i-- {
		if s[i] == '(' {
			l++
		} else {
			r++
		}
		if l == r {
			if 2*r > maxLen {
				maxLen = 2 * l
			}
		} else if l > r {
			l, r = 0, 0
		}
	}
	return maxLen
}
