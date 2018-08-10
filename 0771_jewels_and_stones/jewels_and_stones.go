package leetcode0771

func numJewelsInStones(J string, S string) int {
	m := map[byte]struct{}{}
	for i := range J {
		m[J[i]] = struct{}{}
	}
	res := 0
	for i := range S {
		if _, ok := m[S[i]]; ok {
			res++
		}
	}
	return res
}
