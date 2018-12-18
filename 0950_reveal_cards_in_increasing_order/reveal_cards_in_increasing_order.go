package leetcode0950

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	n := len(deck)
	sort.Ints(deck)
	q := make([]int, n, 2*n)
	for i := 0; i < n; i++ {
		q[i] = i
	}
	index := 0
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[q[index]] = deck[i]
		index++
		if i != n-1 {
			q = append(q, q[index])
			index++
		}
	}
	return res
}
