package leetcode0860

func lemonadeChange(bills []int) bool {
	c5, c10 := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			c5++
		} else if bill == 10 {
			c5--
			c10++
		} else if bill == 20 {
			if c10 > 0 {
				c10--
				c5--
			} else {
				c5 -= 3
			}
		}
		if c5 < 0 || c10 < 0 {
			return false
		}
	}
	return true
}
