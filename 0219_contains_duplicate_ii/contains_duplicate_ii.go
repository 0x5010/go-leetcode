package leetcode0219

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, num := range nums {
		v, ok := m[num]
		if ok && i-v <= k {
			return true
		}
		m[num] = i
	}
	return false
}
