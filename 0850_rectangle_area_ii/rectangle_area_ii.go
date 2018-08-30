package leetcode0850

import "sort"

func rectangleArea(rectangles [][]int) int {
	mod := 1000000007
	n := len(rectangles)
	m := map[int]struct{}{}
	edges := make([]*edge, 2*n)
	for i, r := range rectangles {
		m[r[0]] = struct{}{}
		m[r[2]] = struct{}{}
		edges[2*i] = &edge{x1: r[0], x2: r[2], y: r[1], sig: 1}
		edges[2*i+1] = &edge{x1: r[0], x2: r[2], y: r[3], sig: -1}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].y < edges[j].y
	})
	nx := len(m)
	xs := make([]int, 0, nx)
	for x := range m {
		xs = append(xs, x)
	}
	sort.Ints(xs)
	xindex := make(map[int]int, nx)
	for i, x := range xs {
		xindex[x] = i
	}

	count := make([]int, nx)
	curY, curXSum := 0, 0
	res := 0
	for _, e := range edges {
		res += (e.y - curY) * curXSum
		curY = e.y
		curXSum = 0
		for i := xindex[e.x1]; i < xindex[e.x2]; i++ {
			count[i] += e.sig
		}
		for i := 0; i < nx-1; i++ {
			if count[i] > 0 {
				curXSum += xs[i+1] - xs[i]
			}
		}
	}
	return res % mod
}

type edge struct {
	x1, x2, y, sig int
}
