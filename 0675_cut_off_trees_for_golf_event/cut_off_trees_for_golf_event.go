package leetcode0675

import "sort"

func cutOffTree(forest [][]int) int {
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(forest), len(forest[0])
	trees := make([]*tree, 0, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if forest[i][j] > 1 {
				trees = append(trees, &tree{
					height: forest[i][j],
					x:      i,
					y:      j,
				})
			}
		}
	}
	nn := len(trees)
	sort.Slice(trees, func(i, j int) bool {
		return trees[i].height < trees[j].height
	})

	var bfs func(int, int, int, int) int
	bfs = func(startx, starty, endx, endy int) int {
		f := make([]int, m*n)
		for i := 0; i < m*n; i++ {
			f[i] = -1
		}
		f[startx*n+starty] = 0
		queue := [][2]int{[2]int{startx, starty}}
		for len(queue) != 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur[0] == endx && cur[1] == endy {
				return f[cur[0]*n+cur[1]]
			}
			for _, dir := range dirs {
				x, y := cur[0]+dir[0], cur[1]+dir[1]
				if x < 0 || x >= m || y < 0 || y >= n {
					continue
				}
				if forest[x][y] != 0 && f[x*n+y] == -1 {
					f[x*n+y] = f[cur[0]*n+cur[1]] + 1
					queue = append(queue, [2]int{x, y})
				}
			}
		}
		return -1
	}

	res := 0
	startx, starty := 0, 0
	for i := 0; i < nn; i++ {
		step := bfs(startx, starty, trees[i].x, trees[i].y)
		if step == -1 {
			return -1
		}
		res += step
		startx, starty = trees[i].x, trees[i].y
	}
	return res
}

type tree struct {
	height int
	x, y   int
}
