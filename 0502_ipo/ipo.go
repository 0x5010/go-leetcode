package leetcode0502

import (
	"container/heap"
	"sort"
)

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	n := len(Profits)
	pl := make([]*project, n)
	for i := 0; i < n; i++ {
		pl[i] = &project{
			profit:  Profits[i],
			capital: Capital[i],
		}
	}

	sort.Slice(pl, func(i, j int) bool {
		return pl[i].capital < pl[j].capital
	})

	ph := make(projectHeap, 0, n)
	i := 0
	for k > 0 {
		for i < n && pl[i].capital <= W {
			heap.Push(&ph, pl[i])
			i++
		}

		if len(ph) == 0 {
			break
		}
		W += heap.Pop(&ph).(*project).profit
		k--
	}
	return W
}

type project struct {
	profit, capital int
}

type projectHeap []*project

func (p projectHeap) Len() int           { return len(p) }
func (p projectHeap) Less(i, j int) bool { return p[i].profit > p[j].profit }
func (p projectHeap) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *projectHeap) Push(x interface{}) {
	*p = append(*p, x.(*project))
}

func (h *projectHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
