package leetcode0480

import "container/heap"

func medianSlidingWindow(nums []int, k int) []float64 {
	n := len(nums)
	res := make([]float64, n-k+1)
	w := newWindow(nums, k)

	for i := 0; i+k < n; i++ {
		res[i] = w.median()
		w.update(i, i+k, nums[i+k])
	}
	res[n-k] = w.median()
	return res
}

type window struct {
	k   int
	med int
	m   map[int]*entry
	max *maxHeap
	min *minHeap
}

func newWindow(nums []int, k int) *window {
	max, min := &maxHeap{}, &minHeap{}
	m := make(map[int]*entry, len(nums))

	for i := 0; i < k; i++ {
		e := &entry{
			value: nums[i],
		}

		m[i] = e
		if min.Len() == max.Len() {
			heap.Push(min, e)
		} else {
			heap.Push(max, e)
		}
	}

	w := &window{
		k:   k,
		m:   m,
		max: max,
		min: min,
	}
	w.fix()
	return w
}

func (w *window) update(popIdx, pushIdx, pushValue int) {
	e := w.m[popIdx]
	e.value = pushValue
	w.m[pushIdx] = e

	if e.index < len(w.max.intHeap) {
		heap.Fix(w.max, e.index)
	}
	if e.index < len(w.min.intHeap) {
		heap.Fix(w.min, e.index)
	}

	w.fix()
}

func (w *window) fix() {
	for len(w.max.intHeap) > 0 &&
		w.max.intHeap[0].value > w.min.intHeap[0].value {
		w.max.intHeap[0], w.min.intHeap[0] = w.min.intHeap[0], w.max.intHeap[0]
		heap.Fix(w.max, 0)
		heap.Fix(w.min, 0)
	}
}

func (w *window) median() float64 {
	if w.k%2 == 1 {
		return float64(w.min.intHeap[0].value)
	}
	return float64(w.max.intHeap[0].value+w.min.intHeap[0].value) / 2
}

type minHeap struct {
	intHeap
}

type maxHeap struct {
	intHeap
}

func (h maxHeap) Less(i, j int) bool { return h.intHeap[i].value > h.intHeap[j].value }

type entry struct {
	index int
	value int
}

type intHeap []*entry

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i].value < h[j].value }

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *intHeap) Push(x interface{}) {
	n := len(*h)
	e := x.(*entry)
	e.index = n
	*h = append(*h, e)
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
