package leetcode0295

import (
	"container/heap"
)

type MedianFinder struct {
	l *maxHeap
	r *minHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	l, r := &maxHeap{}, &minHeap{}
	heap.Init(l)
	heap.Init(r)
	return MedianFinder{l, r}
}

func (this *MedianFinder) AddNum(num int) {
	if this.l.Len() == this.r.Len() {
		heap.Push(this.l, num)
	} else {
		heap.Push(this.r, num)
	}

	if this.r.Len() != 0 && this.l.intHeap[0] > this.r.intHeap[0] {
		this.l.intHeap[0], this.r.intHeap[0] = this.r.intHeap[0], this.l.intHeap[0]
		heap.Fix(this.l, 0)
		heap.Fix(this.r, 0)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.l.Len() == this.r.Len() {
		return float64(this.l.intHeap[0]+this.r.intHeap[0]) / 2
	}
	return float64(this.l.intHeap[0])
}

type minHeap struct {
	intHeap
}

type maxHeap struct {
	intHeap
}

func (h maxHeap) Less(i, j int) bool { return h.intHeap[i] > h.intHeap[j] }

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *intHeap) min() int {
	if len(*h) > 0 {
		return (*h)[0]
	}
	return -1
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
