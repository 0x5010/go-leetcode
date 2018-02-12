package leetcode0040

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
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
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			dfs(remain-candidates[i], append(combo, candidates[i]), i+1)
		}
	}
	dfs(target, []int{}, 0)
	return res
}
