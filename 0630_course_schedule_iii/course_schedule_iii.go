package leetcode0630

import (
	"container/heap"
	"sort"
)

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		if courses[i][1] == courses[j][1] {
			return courses[i][0] < courses[j][0]
		}
		return courses[i][1] < courses[j][1]
	})

	res := 0
	sum := 0
	h := &highHeap{}
	for _, course := range courses {
		if course[0]+sum <= course[1] || (h.Len() > 0 && course[0] <= h.Peek()) {
			heap.Push(h, course[0])
			sum += course[0]
			if sum > course[1] {
				sum -= heap.Pop(h).(int)
			}
			res = max(res, h.Len())
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type highHeap []int

func (h highHeap) Len() int           { return len(h) }
func (h highHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h highHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h highHeap) Peek() int {
	if len(h) > 0 {
		return h[0]
	}
	return 0
}

func (h *highHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *highHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return res
}
