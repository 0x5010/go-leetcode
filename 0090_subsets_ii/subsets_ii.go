package leetcode0090

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	var dfs func(int, []int)
	dfs = func(index int, tmp []int) {
		t := make([]int, len(tmp))
		copy(t, tmp)
		res = append(res, t)

		for i := index; i < len(nums); i++ {
			if i == index || nums[i] != nums[i-1] {
				dfs(i+1, append(tmp, nums[i]))
			}
		}
	}

	dfs(0, make([]int, 0, len(nums)))
	return res
}
