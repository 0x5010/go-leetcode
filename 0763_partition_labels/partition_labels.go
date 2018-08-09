package leetcode0763

func partitionLabels(S string) []int {
	n := len(S)
	m := [26]int{}
	for i := 0; i < n; i++ {
		m[S[i]-'a'] = i
	}
	l, r := 0, m[S[0]-'a']
	res := []int{}
	for i := 0; i < n; i++ {
		if i < r {
			r = max(r, m[S[i]-'a'])
			continue
		}
		res = append(res, i-l+1)
		l = i + 1
		if l < len(S) {
			r = m[S[l]-'a']
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
