package leetcode0939

import (
	"math"
	"sort"
)

func minAreaRect(points [][]int) int {
	n := len(points)
	xs, ys := map[int]struct{}{}, map[int]struct{}{}
	for _, point := range points {
		xs[point[0]] = struct{}{}
		ys[point[1]] = struct{}{}
	}
	nx, ny := len(xs), len(ys)
	if nx == n || ny == n {
		return 0
	}
	p := map[int][]int{}
	if nx > ny {
		for _, point := range points {
			p[point[0]] = append(p[point[0]], point[1])
		}
	} else {
		for _, point := range points {
			p[point[1]] = append(p[point[1]], point[0])
		}
		xs, nx = ys, ny
	}
	xl := make([]int, 0, nx)
	for x := range xs {
		xl = append(xl, x)
	}
	sort.Ints(xl)
	lastx := map[[2]int]int{}
	res := math.MaxInt32
	for _, x := range xl {
		yl := p[x]
		sort.Ints(yl)
		nyl := len(yl)
		for i := 0; i < nyl; i++ {
			for j := 0; j < i; j++ {
				y1, y2 := yl[j], yl[i]
				if lx, ok := lastx[[2]int{y1, y2}]; ok {
					res = min(res, (x-lx)*(y2-y1))
				}
				lastx[[2]int{y1, y2}] = x
			}
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
