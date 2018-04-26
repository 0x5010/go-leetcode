package leetcode0347

import "sort"

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		m[num]++
	}

	counts := make([][2]int, 0, len(m))
	for num, count := range m {
		counts = append(counts, [2]int{count, num})
	}
	sortByCount(counts)

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = counts[i][1]
	}
	return res
}

func sortByCount(counts [][2]int) {
	sort.Slice(counts, func(i, j int) bool {
		return counts[i][0] > counts[j][0]
	})
}
