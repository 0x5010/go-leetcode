package leetcode0719

import "sort"

func smallestDistancePair(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)

	count := func(m int) int {
		count, i, j := 0, 0, 1
		for j < n {
			if nums[j]-nums[i] <= m {
				count += j - i
				j++
			} else {
				i++
			}
		}
		return count
	}

	l := nums[1] - nums[0]
	for i := 2; i < n; i++ {
		l = min(l, nums[i]-nums[i-1])
	}

	r := nums[n-1] - nums[0]
	for l < r {
		m := (l + r) / 2
		if count(m) < k {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
