package leetcode0871

import (
	"container/heap"
)

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	n := len(stations)
	gases := intHeap{}
	miles, res := startFuel, 0
	i := 0
	for {
		if miles >= target {
			return res
		}

		for i < n && stations[i][0] <= miles {
			heap.Push(&gases, stations[i][1])
			i++
		}

		if len(gases) == 0 {
			break
		}
		gas := heap.Pop(&gases).(int)
		res++
		miles += gas
	}
	return -1
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] > h[j] }
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
