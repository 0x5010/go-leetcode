package leetcode0685

func findRedundantDirectedConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	var secondLink []int
	for _, e := range edges {
		if parent[e[1]] != 0 {
			secondLink = []int{e[0], e[1]}
		} else {
			parent[e[1]] = e[0]
		}
	}

	if secondLink != nil {
		secondParent := secondLink[0]
		secondParentFound := false
		startNode := secondLink[1]
		node := 0
		for node = parent[startNode]; node != 0 && node != startNode; node = parent[node] {
			if node == secondParent {
				secondParentFound = true
				break
			}
		}
		if node == 0 || secondParentFound {
			return secondLink
		}
		return []int{parent[secondLink[1]], secondLink[1]}
	}
	visited := map[int]struct{}{}
	node := 1
	found := false
	for !found {
		visited[node] = struct{}{}
		node = parent[node]
		_, found = visited[node]
	}
	return []int{parent[node], node}
}
