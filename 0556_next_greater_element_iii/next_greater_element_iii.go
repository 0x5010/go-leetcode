package leetcode0556

import "math"

func nextGreaterElement(n int) int {
	nums := []int{}
	for n != 0 {
		nums = append(nums, n%10)
		n /= 10
	}
	m := len(nums)
	if m < 2 {
		return -1
	}

	i := 1
	for i < m {
		if nums[i] < nums[i-1] {
			break
		}
		i++
	}

	if i == m {
		return -1
	}

	j := 0
	for j < i {
		if nums[j] > nums[i] {
			break
		}
		j++
	}
	nums[i], nums[j] = nums[j], nums[i]
	for k := 0; k < i/2; k++ {
		nums[k], nums[i-k-1] = nums[i-k-1], nums[k]
	}
	res := 0
	p := 1
	for _, n := range nums {
		res += n * p
		p *= 10
	}
	if res > math.MaxInt32 {
		return -1
	}
	return res
}
