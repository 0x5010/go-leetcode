package leetcode0882

func reachableNodes(edges [][]int, M int, N int) int {
	g := make([][][2]int, N)
	for _, e := range edges {
		i, j, n := e[0], e[1], e[2]
		g[i] = append(g[i], [2]int{j, n})
		g[j] = append(g[j], [2]int{i, n})
	}
	queue1, queue2 := []int{0}, []int{}
	moved := make([]int, N)
	moved[0] = 1
	res := 1
	for len(queue1) != 0 {
		for _, i := range queue1 {
			for _, e := range g[i] {
				j, n := e[0], e[1]
				m := moved[i] + n + 1
				if m <= M+1 {
					tmp := moved[j]
					if tmp == 0 || m < tmp {
						moved[j] = m
						queue2 = append(queue2, j)
						if tmp == 0 {
							res++
						}
					}
				}
			}
		}
		queue1, queue2 = queue2, queue1[:0]
	}
	for _, e := range edges {
		i, j, n := e[0], e[1], e[2]
		m := 0
		if moved[i] != 0 {
			m += M + 1 - moved[i]
		}
		if moved[j] != 0 {
			m += M + 1 - moved[j]
		}
		if m > n {
			res += n
		} else {
			res += m
		}
	}
	return res
}
