package leetcode0921

func minAddToMakeValid(S string) int {
	n := len(S)
	l, r := 0, 0
	for i := 0; i < n; i++ {
		if S[i] == '(' {
			l++
		} else if l <= 0 {
			r++
		} else {
			l--
		}
	}
	return l + r
}
