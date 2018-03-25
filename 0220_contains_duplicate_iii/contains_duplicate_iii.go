package leetcode0220

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k < 1 || t < 0 {
		return false
	}
	m := make(map[int]int)
	t++
	for i, num := range nums {
		bucketid := num / t
		if num < 0 {
			bucketid--
		}
		if _, ok := m[bucketid]; ok {
			return true
		} else if v, ok := m[bucketid-1]; ok && abs(num, v) < t {
			return true
		} else if v, ok := m[bucketid+1]; ok && abs(num, v) < t {
			return true
		}
		m[bucketid] = num
		if i >= k {
			delete(m, nums[i-k]/t)
		}
	}
	return false
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
