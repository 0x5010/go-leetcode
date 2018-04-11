package leetcode0310

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = []int{}
	}

	counts := make([]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
		counts[e[0]]++
		counts[e[1]]++
	}

	leaves := make([]int, 0, n)
	for i, c := range counts {
		if c == 1 {
			leaves = append(leaves, i)
		}
	}

	for n > 2 {
		tmp := make([]int, 0, len(leaves))
		for _, leaf := range leaves {
			n--
			for _, node := range g[leaf] {
				counts[node]--
				if counts[node] == 1 {
					tmp = append(tmp, node)
				}
			}
		}
		leaves = tmp
	}
	return leaves
}
