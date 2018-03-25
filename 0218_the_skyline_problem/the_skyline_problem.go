package leetcode0218

import (
	"container/heap"
	"sort"
)

func getSkyline(buildings [][]int) [][]int {
	points := [][]int{}

	highs := new(highHeap)
	heap.Init(highs)
	pre := 0
	heap.Push(highs, pre)

	edges := make([][2]int, 0, 2*len(buildings))
	for _, b := range buildings {
		edges = append(edges, [2]int{b[0], -b[2]})
		edges = append(edges, [2]int{b[1], b[2]})
	}
	sortByEdge(edges)

	for _, e := range edges {
		if e[1] < 0 {
			heap.Push(highs, -1*e[1])
		} else {
			i := 0
			for i < len(*highs) {
				if (*highs)[i] == e[1] {
					break
				}
				i++
			}
			heap.Remove(highs, i)
		}
		now := (*highs)[0]
		if pre != now {
			points = append(points, []int{e[0], now})
			pre = now
		}
	}
	return points
}

type highHeap []int

func (h highHeap) Len() int           { return len(h) }
func (h highHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h highHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *highHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *highHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return res
}

func sortByEdge(edges [][2]int) {
	sort.Slice(edges, func(i, j int) bool {
		a, b := edges[i], edges[j]
		if a[0] == b[0] {
			return a[1] < b[1]
		}
		return a[0] < b[0]
	})
}
