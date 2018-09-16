package leetcode0900

type RLEIterator struct {
	i, n int
	l    []int
}

func Constructor(A []int) RLEIterator {
	return RLEIterator{n: len(A), l: A}
}

func (this *RLEIterator) Next(n int) int {
	for this.i < this.n && n > this.l[this.i] {
		n -= this.l[this.i]
		this.i += 2
	}
	if this.i >= this.n {
		return -1
	}
	this.l[this.i] -= n
	return this.l[this.i+1]
}

/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(A);
 * param_1 := obj.Next(n);
 */
