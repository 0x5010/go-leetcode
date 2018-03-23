package leetcode0210

func findOrder(numCourses int, prerequisites [][]int) []int {
	next := make([][]int, numCourses)
	pre := make([]int, numCourses)
	for _, r := range prerequisites {
		next[r[1]] = append(next[r[1]], r[0])
		pre[r[0]]++
	}

	res := make([]int, numCourses)
	var i, j int
	for i = 0; i < numCourses; i++ {

		for j = 0; j < numCourses; j++ {
			if pre[j] == 0 {
				break
			}
		}
		if j == numCourses {
			return nil
		}

		pre[j] = -1
		for _, c := range next[j] {
			pre[c]--
		}

		res[i] = j
	}
	return res
}
