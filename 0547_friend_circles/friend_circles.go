package leetcode0547

func findCircleNum(M [][]int) int {
	n := len(M)
	if n == 0 {
		return 0
	}
	visited := make([]bool, n)
	res := 0

	var dfs func(int)
	dfs = func(i int) {
		for j := 0; j < n; j++ {
			if M[i][j] == 1 && !visited[j] {
				visited[j] = true
				dfs(j)
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i)
			res++
		}
	}
	return res
}
