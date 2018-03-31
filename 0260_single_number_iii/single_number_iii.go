package leetcode0260

func singleNumber(nums []int) []int {
	diff, a, b := 0, 0, 0
	for _, num := range nums {
		diff ^= num
	}
	diff &= -diff
	for _, num := range nums {
		if num&diff == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
