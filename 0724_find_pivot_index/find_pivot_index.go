package leetcode0724

func pivotIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	l := 0
	for i, num := range nums {
		sum -= num
		if l == sum {
			return i
		}
		l += num
	}
	return -1
}
