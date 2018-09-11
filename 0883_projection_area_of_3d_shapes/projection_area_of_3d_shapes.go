package leetcode0883

func projectionArea(grid [][]int) int {
	n := len(grid)
	xz, yz := make([]int, n), make([]int, n)
	res := 0
	for x, axis := range grid {
		for y, z := range axis {
			if z == 0 {
				continue
			}
			res++
			xz[y] = max(xz[y], z)
			yz[x] = max(yz[x], z)
		}
	}
	for i := range xz {
		res += xz[i] + yz[i]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
