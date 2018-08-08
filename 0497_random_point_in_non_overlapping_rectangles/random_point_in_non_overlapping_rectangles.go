package leetcode0497

import (
	"math/rand"
	"sort"
)

type Solution struct {
	counts []int
	rects  []*rect
}

type rect struct {
	x, y, w, h int
}

func Constructor(rects [][]int) Solution {
	n := len(rects)
	total := 0
	counts := make([]int, n)
	rs := make([]*rect, n)
	for i, r := range rects {
		w, h := r[2]-r[0]+1, r[3]-r[1]+1
		total += w * h
		counts[i] = total
		rs[i] = &rect{
			x: r[0],
			y: r[1],
			w: w,
			h: h,
		}
	}
	return Solution{
		counts: counts,
		rects:  rs,
	}
}

func (this *Solution) Pick() []int {
	cand := rand.Intn(this.counts[len(this.counts)-1]) + 1

	i := sort.SearchInts(this.counts, cand)
	r := this.rects[i]
	return []int{
		r.x + rand.Intn(r.w),
		r.y + rand.Intn(r.h),
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */
