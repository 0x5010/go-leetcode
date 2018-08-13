package leetcode0785

func isBipartite(graph [][]int) bool {
	colors := make([]int, len(graph))

	var color func(int) bool
	color = func(i int) bool {
		for _, j := range graph[i] {
			if colors[j] == 0 {
				colors[j] = -colors[i]
				if !color(j) {
					return false
				}
			} else if colors[i] == colors[j] {
				return false
			}
		}
		return true
	}
	for i := range graph {
		if colors[i] == 0 {
			colors[i] = 1
			if !color(i) {
				return false
			}
		}
	}
	return true
}
