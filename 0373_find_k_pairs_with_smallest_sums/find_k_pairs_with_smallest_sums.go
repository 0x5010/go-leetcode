package leetcode0373

import "container/heap"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	m, n := len(nums1), len(nums2)
	if m == 0 || n == 0 {
		return nil
	}

	psLen := min(k, m)
	ps := make(pairs, psLen)
	for i := 0; i < psLen; i++ {
		ps[i] = &pair{i, 0, nums1[i] + nums2[0]}
	}

	heap.Init(&ps)

	res := [][]int{}
	for k > 0 && len(ps) > 0 {
		min := heap.Pop(&ps).(*pair)
		res = append(res, []int{nums1[min.i], nums2[min.j]})
		if min.j+1 < n {
			heap.Push(&ps, &pair{min.i, min.j + 1, nums1[min.i] + nums2[min.j+1]})
		}
		k--
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type pair struct {
	i, j int
	sum  int
}

type pairs []*pair

func (ps pairs) Len() int           { return len(ps) }
func (ps pairs) Less(i, j int) bool { return ps[i].sum < ps[j].sum }
func (ps pairs) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }

func (ps *pairs) Push(x interface{}) {
	p := x.(*pair)
	*ps = append(*ps, p)
}

func (ps *pairs) Pop() interface{} {
	old := *ps
	n := len(old)
	p := old[n-1]
	*ps = old[:n-1]
	return p
}
