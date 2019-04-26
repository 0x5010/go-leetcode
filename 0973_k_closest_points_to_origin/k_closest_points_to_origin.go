package leetcode0973

func kClosest(points [][]int, K int) [][]int {
	dist := func(i int) int {
		x, y := points[i][0], points[i][1]
		return x*x + y*y
	}
	partition := func(l, r int) int {
		pivot, pdist := l, dist(l)

		points[pivot], points[r] = points[r], points[pivot]
		divider := l
		for i := l; i < r; i++ {
			if dist(i) < pdist {
				points[i], points[divider] = points[divider], points[i]
				divider++
			}
		}
		points[divider], points[r] = points[r], points[divider]
		return divider
	}
	l, r := 0, len(points)-1
	for l < r {
		i := partition(l, r)
		if i == K {
			break
		} else if i > K {
			r = i - 1
		} else {
			l = i + 1
		}
	}
	return points[:K]
}
