package leetcode0491

// TLE
func findSubsequences(nums []int) [][]int {
	res := [][]int{}
	seq := []int{}

	var dfs func(int)
	dfs = func(index int) {
		n := len(seq)
		if n > 1 {
			tmp := make([]int, n)
			copy(tmp, seq)
			res = append(res, tmp)
		}
		m := make(map[int]struct{})
		for i := index; i < len(nums); i++ {
			if _, ok := m[nums[i]]; ok {
				continue
			}
			if index == len(nums) {
				return
			}
			m[nums[i]] = struct{}{}
			if n == 0 || nums[i] >= seq[n-1] {
				seq = append(seq, nums[i])
				dfs(i + 1)
				seq = seq[:n]
			}
		}
	}

	dfs(0)
	return res
}
