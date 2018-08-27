package leetcode0834

func sumOfDistancesInTree(N int, edges [][]int) []int {
	tree := make([][]int, N)
	for _, e := range edges {
		i, j := e[0], e[1]
		tree[i] = append(tree[i], j)
		tree[j] = append(tree[j], i)
	}

	count := make([]int, N)
	res := make([]int, N)

	visited := make([]bool, N)
	var dfs func(int)
	dfs = func(i int) {
		visited[i] = true
		for _, n := range tree[i] {
			if visited[n] {
				continue
			}
			dfs(n)
			count[i] += count[n]
			res[i] += res[n] + count[n]
		}
		count[i]++
	}
	dfs(0)

	visited = make([]bool, N)
	var dfs2 func(int)
	dfs2 = func(i int) {
		visited[i] = true
		for _, n := range tree[i] {
			if visited[n] {
				continue
			}
			res[n] = res[i] + (N - count[n]) - count[n]
			dfs2(n)
		}
	}
	dfs2(0)

	return res
}
