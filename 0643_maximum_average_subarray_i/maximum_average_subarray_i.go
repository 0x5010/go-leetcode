package leetcode0643

func findMaxAverage(nums []int, k int) float64 {
	tmp := 0
	for i := 0; i < k; i++ {
		tmp += nums[i]
	}

	max := tmp
	for i := k; i < len(nums); i++ {
		tmp = tmp - nums[i-k] + nums[i]

		if max < tmp {
			max = tmp
		}
	}

	return float64(max) / float64(k)
}
