package leetcode0699

import "container/heap"

func fallingSquares(positions [][]int) []int {
	n := len(positions)
	res := make([]int, n)
	h := make(intervals, 0, n)

	for i := 0; i < n; i++ {
		sp := &interval{
			l: positions[i][0],
			r: positions[i][0] + positions[i][1],
			h: positions[i][1],
		}

		height := 0
		removes := make(intervals, 0, len(h))

		for j := 0; j < len(h); j++ {
			if h[j].r <= sp.l || sp.r <= h[j].l {
				continue
			}

			height = max(height, h[j].h)

			if sp.l <= h[j].l && h[j].r <= sp.r {
				removes = append(removes, h[j])
			}

			if h[j].l < sp.l && sp.r < h[j].r {
				heap.Push(&h, &interval{
					l: sp.r,
					r: h[j].r,
					h: h[j].h,
				})
				h[j].r = sp.l
				break
			}

			if h[j].l < sp.l && sp.l < h[j].r && h[j].r <= sp.r {
				h[j].r = sp.l
			}

			if sp.l <= h[j].l && h[j].l < sp.r && sp.r < h[j].r {
				h[j].l = sp.r
			}
		}

		for j := 0; j < len(removes); j++ {
			heap.Remove(&h, removes[j].index)
		}

		sp.h += height
		heap.Push(&h, sp)
		res[i] = h[0].h
	}

	return res
}

type interval struct {
	l, r, h int
	index   int
}
type intervals []*interval

func (is intervals) Len() int { return len(is) }

func (is intervals) Less(i, j int) bool { return is[i].h > is[j].h }

func (is intervals) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
	is[i].index = i
	is[j].index = j
}

func (is *intervals) Push(x interface{}) {
	e := x.(*interval)
	e.index = len(*is)
	*is = append(*is, e)
}

func (is *intervals) Pop() interface{} {
	e := (*is)[len(*is)-1]
	*is = (*is)[0 : len(*is)-1]
	return e
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
