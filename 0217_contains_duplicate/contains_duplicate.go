package leetcode0217

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = struct{}{}
	}
	return false
}
