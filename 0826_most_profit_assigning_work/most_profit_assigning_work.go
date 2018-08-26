package leetcode0826

import (
	"sort"
)

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
	jobs := make([][2]int, n)
	for i, d := range difficulty {
		jobs[i] = [2]int{d, profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})
	sort.Ints(worker)
	i, maxp := 0, 0
	res := 0
	for _, ability := range worker {
		for i < n && ability >= jobs[i][0] {
			maxp = max(maxp, jobs[i][1])
			i++
		}
		res += maxp
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
