package leetcode0802

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	if n == 0 {
		return nil
	}
	res := []int{}
	color := make([]int, n)
	var dfs func(int) bool
	dfs = func(start int) bool {
		if color[start] != 0 {
			return color[start] == 1
		}
		color[start] = 2
		for _, node := range graph[start] {
			if !dfs(node) {
				return false
			}
		}
		color[start] = 1
		return true
	}
	for i := 0; i < n; i++ {
		if dfs(i) {
			res = append(res, i)
		}
	}
	return res
}
