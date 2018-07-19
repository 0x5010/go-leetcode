package leetcode0621

func leastInterval(tasks []byte, n int) int {
	c := make([]int, 26)
	for _, task := range tasks {
		c[task-'A']++
	}

	most := 0
	for _, count := range c {
		if most < count {
			most = count
		}
	}
	idles := (most - 1) * (n + 1)
	for _, count := range c {
		idles -= min(most-1, count)
	}
	return len(tasks) + max(0, idles)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
