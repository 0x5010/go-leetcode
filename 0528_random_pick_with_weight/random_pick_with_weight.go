package leetcode0528

import (
	"math/rand"
	"sort"
)

type Solution struct {
	w     []int
	total int
}

func Constructor(w []int) Solution {
	total := 0
	ws := make([]int, len(w))
	for i, v := range w {
		total += v
		ws[i] = total
	}
	return Solution{
		w:     ws,
		total: total,
	}
}

func (this *Solution) PickIndex() int {
	r := rand.Intn(this.total)
	return sort.Search(len(this.w), func(i int) bool {
		return this.w[i] > r
	})
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
