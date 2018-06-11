package leetcode0452

import "sort"

func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n == 0 {
		return 0
	}
	sortByEnd(points)

	res := 1
	end := points[0][1]

	for i := 1; i < n; i++ {
		if points[i][0] <= end {
			continue
		}
		res++
		end = points[i][1]
	}
	return res
}

func sortByEnd(points [][]int) {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
}
