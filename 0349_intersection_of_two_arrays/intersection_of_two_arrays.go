package leetcode0349

func intersection(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	if n1 > n2 {
		n1, n2, nums1, nums2 = n2, n1, nums2, nums1
	}
	m := make(map[int]struct{}, n1)
	for _, num := range nums1 {
		m[num] = struct{}{}
	}

	res := make([]int, 0, n1)
	for _, num := range nums2 {
		if _, ok := m[num]; ok {
			res = append(res, num)
			delete(m, num)
		}
	}
	return res
}
