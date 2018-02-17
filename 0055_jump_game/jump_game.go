package leetcode0055

func canJump(nums []int) bool {
	max := 0
	for i, v := range nums {
		if v > 0 {
			if i+v > max {
				max = i + v
			}
		} else {
			if max == i {
				if i == len(nums)-1 {
					return true
				}
				return false
			}
		}

	}
	return max >= len(nums)-1
}
