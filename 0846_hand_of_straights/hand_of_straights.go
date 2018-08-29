package leetcode0846

import "container/heap"

func isNStraightHand(hand []int, W int) bool {
	h := intHeap(hand)
	heap.Init(&h)
	for h.Len() > 0 {
		start := heap.Pop(&h).(int)
		for i := 1; i < W; i++ {
			if !h.remove(start + i) {
				return false
			}
		}
	}
	return true
}

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

func (h *intHeap) remove(x interface{}) bool {
	for i := 0; i < h.Len(); i++ {
		if (*h)[i] == x.(int) {
			heap.Remove(h, i)
			return true
		}
	}
	return false
}
