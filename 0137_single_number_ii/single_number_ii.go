package leetcode0137

func singleNumber(nums []int) int {
	a, b := 0, 0
	for _, n := range nums {
		a = (a ^ n) & ^b
		b = (b ^ n) & ^a
	}
	return a
}
