package leetcode0974

func subarraysDivByK(A []int, K int) int {
	count := make([]int, K)
	count[0] = 1
	p, res := 0, 0
	for _, a := range A {
		p = (p + a) % K
		if p < 0 {
			p += K
		}
		res += count[p]
		count[p]++
	}
	return res
}
