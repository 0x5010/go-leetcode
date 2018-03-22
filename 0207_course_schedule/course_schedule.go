package leetcode0207

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	inEdges := make([]int, numCourses)
	res := []int{}
	queue := make([]int, 0)

	for i := 0; i < len(prerequisites); i++ {
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
		inEdges[prerequisites[i][0]]++
	}

	for i := 0; i < numCourses; i++ {
		if inEdges[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		course := queue[0]
		res = append(res, course)
		queue = queue[1:]

		for i := 0; i < len(graph[course]); i++ {
			inEdges[graph[course][i]]--

			if inEdges[graph[course][i]] == 0 {
				queue = append(queue, graph[course][i])
			}
		}
	}
	return len(res) == numCourses
}
