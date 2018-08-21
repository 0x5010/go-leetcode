package leetcode0815

func numBusesToDestination(routes [][]int, S int, T int) int {
	if S == T {
		return 0
	}

	hasT := false
	buses := make(map[int][]int, len(routes))
	for i, route := range routes {
		for _, r := range route {
			buses[r] = append(buses[r], i)
			if r == T {
				hasT = true
			}
		}
	}
	if !hasT {
		return -1
	}

	visited := map[int]struct{}{}
	q := []int{S}
	res := 0
	for len(q) != 0 {
		n := len(q)
		res++
		for i := 0; i < n; i++ {
			cur := q[i]
			bs := buses[cur]
			for _, bus := range bs {
				if _, ok := visited[bus]; ok {
					continue
				}
				visited[bus] = struct{}{}
				for _, r := range routes[bus] {
					if r == T {
						return res
					}
					q = append(q, r)
				}
			}
		}
		q = q[n:]
	}
	return -1
}
