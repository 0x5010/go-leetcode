package leetcode0523

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	m := make(map[int]int, n)
	m[0] = -1

	sum := 0
	for i, num := range nums {
		sum += num
		if k != 0 {
			sum %= k
		}

		if index, ok := m[sum]; ok {
			if i-index > 1 {
				return true
			}
		} else {
			m[sum] = i
		}
	}
	return false
}
