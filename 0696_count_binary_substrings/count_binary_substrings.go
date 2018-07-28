package leetcode0696

func countBinarySubstrings(s string) int {
	prev, cur := 0, 1
	res := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cur++
		} else {
			prev = cur
			cur = 1
		}
		if prev >= cur {
			res++
		}
	}
	return res
}
