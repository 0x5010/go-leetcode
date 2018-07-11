package leetcode0554

func leastBricks(wall [][]int) int {
	n := len(wall)
	if n == 0 {
		return 0
	}
	m := map[int]int{}
	count := 0
	for _, row := range wall {
		l := 0
		for i := 0; i < len(row)-1; i++ {
			l += row[i]
			m[l]++
			count = max(count, m[l])
		}
	}
	return n - count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
