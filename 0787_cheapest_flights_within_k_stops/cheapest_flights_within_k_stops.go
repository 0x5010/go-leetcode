package leetcode0787

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	g := make([][][2]int, n)
	for _, f := range flights {
		g[f[0]] = append(g[f[0]], [2]int{f[1], f[2]})
	}
	p1, p2 := make([]int, n), make([]int, n)
	q1, q2 := []int{src}, []int{}
	p2[src] = 1
	for len(q1) != 0 && K >= 0 {
		for _, f := range q1 {
			p1[f] = p2[f]
		}
		for _, f := range q1 {
			p := p1[f]
			for _, nn := range g[f] {
				price := p + nn[1]
				dst := nn[0]
				if p2[dst] == 0 || p2[dst] > price {
					p2[dst] = price
					q2 = append(q2, dst)
				}
			}
		}
		q1, q2 = q2, q1[:0]
		K--
	}
	return p2[dst] - 1
}
