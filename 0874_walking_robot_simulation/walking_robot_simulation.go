package leetcode0874

func robotSim(commands []int, obstacles [][]int) int {
	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	m := map[obstacle]struct{}{}
	for _, o := range obstacles {
		m[obstacle{o[0], o[1]}] = struct{}{}
	}
	i, j, d := 0, 0, 0
	res := 0
	for _, c := range commands {
		if c == -2 {
			d = (d + 3) % 4
		} else if c == -1 {
			d = (d + 1) % 4
		} else {
			dx, dy := dirs[d][0], dirs[d][1]
			for k := 0; k < c; k++ {
				x, y := i+dx, j+dy
				if _, ok := m[obstacle{x, y}]; ok {
					break
				}
				i = x
				j = y
			}
			res = max(res, i*i+j*j)
		}
	}
	return res
}

type obstacle struct {
	x, y int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
