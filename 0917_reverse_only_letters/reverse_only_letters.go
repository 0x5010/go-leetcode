package leetcode0917

func reverseOnlyLetters(S string) string {
	n := len(S)
	bs := []byte(S)
	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		for l < r && !isChar(bs[l]) {
			l++
		}
		for l < r && !isChar(bs[r]) {
			r--
		}
		if l < r {
			bs[l], bs[r] = bs[r], bs[l]
		}
	}
	return string(bs)
}

func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') {
		return true
	}
	return false
}
