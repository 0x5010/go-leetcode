package leetcode0005

func longestPalindrome(s string) string {
	n := len(s)
	max, start := 0, 0
	for i := 0; i < n; {
		l, r := i, i
		for r < (n-1) && s[r] == s[r+1] {
			r++
		}
		i = r + 1
		for l > 0 && r < n-1 && s[l-1] == s[r+1] {
			l--
			r++
		}
		tmp := r - l + 1
		if tmp > max {
			max = tmp
			start = l
		}
	}
	return s[start : start+max]
}
