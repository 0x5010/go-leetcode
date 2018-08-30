package leetcode0849

func maxDistToClosest(seats []int) int {
	e := 0
	res := 0
	for i, s := range seats {
		if e == i {
			res = e
		} else {
			res = max(res, (e+e%2)/2)
		}
		if s == 1 {
			e = 0
		} else {
			e++
		}
	}
	return max(res, e)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
