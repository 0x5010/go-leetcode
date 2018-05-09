package leetcode0391

import "math"

func isRectangleCover(rectangles [][]int) bool {
	if len(rectangles) == 0 || len(rectangles[0]) == 0 {
		return false
	}
	x1, x2, y1, y2 := math.MaxInt32, math.MinInt32, math.MaxInt32, math.MinInt32

	type point struct {
		x, y int
	}
	m := make(map[point]struct{})
	area := 0

	for _, r := range rectangles {
		x1 = min(r[0], x1)
		y1 = min(r[1], y1)
		x2 = max(r[2], x2)
		y2 = max(r[3], y2)

		area += (r[2] - r[0]) * (r[3] - r[1])

		for _, p := range []point{
			point{r[0], r[1]},
			point{r[0], r[3]},
			point{r[2], r[1]},
			point{r[2], r[3]},
		} {
			if _, ok := m[p]; ok {
				delete(m, p)
			} else {
				m[p] = struct{}{}
			}
		}
	}
	if _, ok := m[point{x1, y1}]; !ok {
		return false
	}
	if _, ok := m[point{x1, y2}]; !ok {
		return false
	}
	if _, ok := m[point{x2, y1}]; !ok {
		return false
	}
	if _, ok := m[point{x2, y2}]; !ok {
		return false
	}
	if len(m) != 4 {
		return false
	}
	return area == (x2-x1)*(y2-y1)
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
