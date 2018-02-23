package leetcode0075

func sortColors(nums []int) {
	n := len(nums)

	if n == 0 {
		return
	}

	tmp := nums[0]
	nums[0] = 1

	i, j, k := 0, 1, n-1
	for j <= k {
		if nums[j] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else if nums[j] == 1 {
			j++
		} else {
			nums[j], nums[k] = nums[k], nums[j]
			k--
		}
	}

	if tmp == 0 {
		nums[i] = tmp
	} else if tmp == 2 {
		nums[k] = tmp
	}
}
