package leetcode0812

import "math"

func largestTriangleArea(points [][]int) float64 {
	res := 0.
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				res = max(res, area(points[i], points[j], points[k]))
			}
		}
	}
	return res
}

func area(p1, p2, p3 []int) float64 {
	return math.Abs(float64(p1[0]*p2[1]+p2[0]*p3[1]+p3[0]*p1[1]-p1[0]*p3[1]-p2[0]*p1[1]-p3[0]*p2[1])) * 0.5
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
