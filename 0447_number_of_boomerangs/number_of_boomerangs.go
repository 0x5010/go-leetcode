package leetcode0447

func numberOfBoomerangs(points [][]int) int {
	n := len(points)
	if n < 3 {
		return 0
	}
	res := 0
	for i := 0; i < n; i++ {
		m := make(map[int]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			x := points[j][0] - points[i][0]
			y := points[j][1] - points[i][1]
			dist := x*x + y*y
			m[dist]++
		}
		for _, v := range m {
			res += v * (v - 1)
		}
	}
	return res
}
