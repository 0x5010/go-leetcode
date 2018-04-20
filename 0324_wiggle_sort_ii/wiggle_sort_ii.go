package leetcode0324

import "sort"

func wiggleSort(nums []int) {
	n := len(nums)
	sort.Ints(nums)
	med := nums[n/2]

	i, l, r := 0, 0, n-1
	for i <= r {
		index := mapIndex(i, n)
		if nums[index] > med {
			swap(nums, mapIndex(l, n), index)
			l++
			i++
		} else if nums[index] < med {
			swap(nums, mapIndex(r, n), index)
			r--
		} else {
			i++
		}
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func mapIndex(i, n int) int {
	return (1 + 2*i) % (n | 1)
}
