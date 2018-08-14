package leetcode0789

func escapeGhosts(ghosts [][]int, target []int) bool {
	max := abs(target[0]) + abs(target[1])
	for _, g := range ghosts {
		d := abs(g[0]-target[0]) + abs(g[1]-target[1])
		if d <= max {
			return false
		}
	}
	return true
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
