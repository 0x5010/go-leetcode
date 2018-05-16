package leetcode0417

func pacificAtlantic(matrix [][]int) [][]int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}
	res := [][]int{}
	p, a := make([][]bool, m), make([][]bool, m)
	for i := 0; i < m; i++ {
		p[i] = make([]bool, n)
		a[i] = make([]bool, n)
	}

	pq, aq := make([][]int, m+n), make([][]int, m+n)
	for i := 0; i < m; i++ {
		p[i][0] = true
		pq[i] = []int{i, 0}
		a[i][n-1] = true
		aq[i] = []int{i, n - 1}
	}
	for j := 0; j < n; j++ {
		p[0][j] = true
		pq[m+j] = []int{0, j}
		a[m-1][j] = true
		aq[m+j] = []int{m - 1, j}
	}
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	bfs := func(q [][]int, b [][]bool) {
		for len(q) > 0 {
			c := q[0]
			q = q[1:]
			for _, dir := range dirs {
				i, j := c[0]+dir[0], c[1]+dir[1]
				if 0 <= i && i < m && 0 <= j && j < n && !b[i][j] && matrix[c[0]][c[1]] <= matrix[i][j] {
					b[i][j] = true
					q = append(q, []int{i, j})
				}
			}
		}
	}
	bfs(pq, p)
	bfs(aq, a)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[i][j] && a[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}
