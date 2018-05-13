package leetcode0409

func longestPalindrome(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	odd := 0
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for _, b := range m {
		count += b / 2
		if b%2 != 0 {
			odd = 1
		}
	}
	return 2*count + odd
}
