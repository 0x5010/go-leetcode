package leetcode0058

func lengthOfLastWord(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	res := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if res != 0 {
				return res
			}
			continue
		}
		res++
	}
	return res
}
