package leetcode0797

func allPathsSourceTarget(graph [][]int) [][]int {
	dst := len(graph) - 1
	path := make([]int, len(graph))
	res := [][]int{}
	var dfs func(int, int)
	dfs = func(id, level int) {
		if id == dst {
			tmp := make([]int, level)
			copy(tmp, path)
			res = append(res, tmp)
		}
		for _, node := range graph[id] {
			path[level] = node
			dfs(node, level+1)
		}
	}
	dfs(0, 1)
	return res
}
