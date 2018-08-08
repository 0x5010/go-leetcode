package leetcode0519

import "math/rand"

type Solution struct {
	m           map[int]int
	r, c, total int
}

func Constructor(n_rows int, n_cols int) Solution {
	return Solution{
		m:     map[int]int{},
		r:     n_rows,
		c:     n_cols,
		total: n_rows * n_cols,
	}
}

func (this *Solution) Flip() []int {
	i := rand.Intn(this.total)
	this.total--
	x := i
	if v, ok := this.m[i]; ok {
		x = v
	}
	if v, ok := this.m[this.total]; ok {
		this.m[i] = v
	} else {
		this.m[i] = this.total
	}
	return []int{x / this.c, x % this.c}
}

func (this *Solution) Reset() {
	this.m = map[int]int{}
	this.total = this.r * this.c
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n_rows, n_cols);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
