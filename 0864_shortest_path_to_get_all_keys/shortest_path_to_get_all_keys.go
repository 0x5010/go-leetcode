package leetcode0864

func shortestPathAllKeys(grid []string) int {
	m, n := len(grid), len(grid[0])
	keys := 0
	visited := map[entry]struct{}{}
	queue := []*entry{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			b := grid[i][j]
			if b >= 'a' {
				keys |= 1 << uint(b-'a')
			} else if b == '@' {
				e := &entry{
					x: i,
					y: j,
				}
				visited[*e] = struct{}{}
				queue = append(queue, e)
			}
		}
	}
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	res := 1
	for len(queue) != 0 {
		count := len(queue)
		for i := 0; i < count; i++ {
			e := queue[i]
			for _, dir := range dirs {
				x, y := e.x+dir[0], e.y+dir[1]
				if x < 0 || x >= m || y < 0 || y >= n {
					continue
				}
				b := grid[x][y]
				if b == '#' ||
					'A' <= b && b <= 'F' && e.keys&(1<<uint(b-'A')) == 0 {
					continue
				}
				ks := e.keys
				if b >= 'a' {
					ks |= 1 << uint(b-'a')
					if ks == keys {
						return res
					}
				}
				next := &entry{
					x:    x,
					y:    y,
					keys: ks,
				}
				if _, ok := visited[*next]; ok {
					continue
				}
				visited[*next] = struct{}{}
				queue = append(queue, next)
			}
		}
		queue = queue[count:]
		res++
	}
	return -1
}

type entry struct {
	x, y, keys int
}
