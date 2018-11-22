package leetcode0942

func diStringMatch(S string) []int {
	n := len(S)
	l, r := 0, n
	res := make([]int, n+1)
	for i := 0; i < n; i++ {
		if S[i] == 'I' {
			res[i] = l
			l++
		} else {
			res[i] = r
			r--
		}
	}
	res[n] = l
	return res
}
