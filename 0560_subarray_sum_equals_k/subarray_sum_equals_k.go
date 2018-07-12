package leetcode0560

func subarraySum(nums []int, k int) int {
	res, sum := 0, 0
	m := map[int]int{}
	m[0] = 1

	for _, num := range nums {
		sum += num
		if v, ok := m[sum-k]; ok {
			res += v
		}
		m[sum]++
	}
	return res
}
