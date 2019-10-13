package leetcode0031

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	reverse := func(start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}

	p := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			p = i
			break
		}
	}

	if p == -1 {
		reverse(0, len(nums)-1)
		return
	}
	nextP := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > nums[p] {
			nextP = i
			break
		}
	}
	nums[p], nums[nextP] = nums[nextP], nums[p]

	reverse(p+1, len(nums)-1)
}
