package leetcode0322

import "sort"

func wiggleSort(nums []int) {
	n := len(nums)
	tmp := make([]int, n)
	copy(tmp, nums)
	sort.Ints(tmp)

	mid := n / 2

	for i := 0; i < n; i++ {
		if i < mid {
			nums[2*i+1] = tmp[n-1-i]
		} else {
			nums[2*(i-mid)] = tmp[n-1-i]
		}
	}
}
