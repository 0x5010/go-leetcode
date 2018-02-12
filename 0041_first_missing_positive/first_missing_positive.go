package leetcode0041

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func firstMissingPositive(nums []int) int {
	i, n := 0, len(nums)

	for i < n {
		if nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			swap(nums, i, nums[i]-1)
		} else {
			i++
		}
	}

	for k := 0; k < n; k++ {
		if nums[k] != k+1 {
			return k + 1
		}
	}
	return n + 1
}
