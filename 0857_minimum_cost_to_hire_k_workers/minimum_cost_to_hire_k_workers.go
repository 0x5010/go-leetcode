package leetcode0857

import (
	"container/heap"
	"math"
	"sort"
)

func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
	n := len(quality)

	workers := make([][2]float64, n)
	for i := 0; i < n; i++ {
		w, q := float64(wage[i]), float64(quality[i])
		workers[i][0], workers[i][1] = w/q, q
	}
	sort.Slice(workers, func(i, j int) bool {
		return workers[i][0] < workers[j][0]
	})

	res := math.MaxFloat64
	sum := 0.

	h := floatHeap{}

	for _, w := range workers {
		sum += w[1]
		heap.Push(&h, w[1])

		if len(h) > K {
			sum -= heap.Pop(&h).(float64)
		}

		if len(h) == K {
			res = min(res, sum*w[0])
		}
	}
	return res
}

type floatHeap []float64

func (h floatHeap) Len() int           { return len(h) }
func (h floatHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h floatHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *floatHeap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}

func (h *floatHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
