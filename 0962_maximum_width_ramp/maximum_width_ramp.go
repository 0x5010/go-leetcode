package leetcode0962

func maxWidthRamp(A []int) int {
	n := len(A)
	s := []int{}
	res := 0
	for i, a := range A {
		if len(s) == 0 || A[s[len(s)-1]] > a {
			s = append(s, i)
		}
	}
	m := len(s) - 1
	for i := n - 1; i > res; i-- {
		for m >= 0 && A[s[m]] <= A[i] {
			res = max(res, i-s[m])
			m--
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
