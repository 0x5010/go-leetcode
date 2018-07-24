package leetcode0659

func isPossible(nums []int) bool {
	c1, c2, c3, p1, p2, p3 := 1, 0, 0, 0, 0, 0
	f := func() {
		if p1 > 0 {
			p1--
			c2++
		} else if p2 > 0 {
			p2--
			c3++
		} else if p3 > 0 {
			p3--
			c3++
		} else {
			c1++
		}
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			f()
		} else if nums[i] == nums[i-1]+1 {
			if p1 > 0 || p2 > 0 {
				return false
			}
			p1, p2, p3 = c1, c2, c3
			c1, c2, c3 = 0, 0, 0
			f()
		} else {
			if c1 > 0 || c2 > 0 {
				return false
			}
			c1, c2, c3, p1, p2, p3 = 1, 0, 0, 0, 0, 0
		}
	}
	return c1 == 0 && c2 == 0
}
