package leetcode0799

func champagneTower(poured int, query_row int, query_glass int) float64 {
	if query_glass > query_row/2 {
		query_glass = query_row - query_glass
	}
	row := [100]float64{0: float64(poured)}
	for i := 0; ; i++ {
		next := [100]float64{}
		for j := max(i+query_glass-query_row, 0); j <= query_glass; j++ {
			if row[j] > 1 {
				e := (row[j] - 1) / 2
				next[j] += e
				next[j+1] += e
				row[j] = 1
			}
		}
		if i == query_row {
			return row[query_glass]
		}
		row = next
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
