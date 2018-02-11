package leetcode0039

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	var dfs func(remain int, combo []int, index int)
	dfs = func(remain int, combo []int, index int) {
		if remain == 0 {
			t := make([]int, len(combo))
			copy(t, combo)
			res = append(res, t)
			return
		}

		for i := index; i < len(candidates); i++ {
			if candidates[i] > remain {
				break
			}
			dfs(remain-candidates[i], append(combo, candidates[i]), i)
		}

	}
	dfs(target, []int{}, 0)
	return res
}
