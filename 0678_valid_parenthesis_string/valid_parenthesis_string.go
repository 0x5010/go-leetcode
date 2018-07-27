package leetcode0678

func checkValidString(s string) bool {
	l, r := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			l++
			r++
		} else if s[i] == ')' {
			if l > 0 {
				l--
			}
			r--
		} else {
			if l > 0 {
				l--
			}
			r++
		}
		if r < 0 {
			return false
		}
	}
	return l == 0
}
