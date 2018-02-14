package leetcode0047

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	taken := make([]bool, n)
	res := [][]int{}

	var makePermutation func(cur int, vector []int)
	makePermutation = func(cur int, vector []int) {
		if cur == n {
			tmp := make([]int, n)
			copy(tmp, vector)
			res = append(res, tmp)
			return
		}
		used := make(map[int]bool, n-cur)
		for i := 0; i < n; i++ {
			if !taken[i] && !used[nums[i]] {
				used[nums[i]] = true
				taken[i] = true
				vector[cur] = nums[i]
				makePermutation(cur+1, vector)
				taken[i] = false
			}
		}
	}

	makePermutation(0, make([]int, n))
	return res
}
