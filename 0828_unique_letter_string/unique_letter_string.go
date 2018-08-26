package leetcode0828

func uniqueLetterString(S string) int {
	mod := 1000000007
	n := len(S)
	index := [26][2]int{}
	for i := 0; i < 26; i++ {
		index[i][0], index[i][1] = -1, -1
	}
	res := 0
	for i := 0; i < n; i++ {
		c := int(S[i] - 'A')
		res += (i - index[c][1]) * (index[c][1] - index[c][0])
		index[c][0], index[c][1] = index[c][1], i
	}
	for i := 0; i < 26; i++ {
		res += (n - index[i][1]) * (index[i][1] - index[i][0])
	}
	return res % mod
}
