package leetcode0703

import (
	"container/heap"
)

type KthLargest struct {
	h intHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	var h intHeap
	if len(nums) < k {
		h = intHeap(nums)
		heap.Init(&h)
	} else {
		h = intHeap(nums[:k])
		heap.Init(&h)
		for i := k; i < len(nums); i++ {
			if nums[i] > h.min() {
				heap.Pop(&h)
				heap.Push(&h, nums[i])
			}
		}
	}
	return KthLargest{
		k: k,
		h: h,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.h, val)
	if this.h.Len() > this.k {
		heap.Pop(&this.h)
	}
	return this.h.min()
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

func (h *intHeap) min() int {
	if len(*h) > 0 {
		return (*h)[0]
	}
	return -1
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
