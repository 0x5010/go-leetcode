package leetcode0647

func countSubstrings(s string) int {
	n := len(s)
	countPalindrome := func(l, r int) int {
		count := 0
		for l >= 0 && r < n && s[l] == s[r] {
			count++
			l--
			r++
		}
		return count
	}

	res := 0
	for i := 0; i < n; i++ {
		res += countPalindrome(i, i) + countPalindrome(i, i+1)
	}
	return res
}
