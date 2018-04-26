package leetcode0350

func intersect(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	if n1 > n2 {
		n1, n2, nums1, nums2 = n2, n1, nums2, nums1
	}
	m := make(map[int]int, n1)
	for _, num := range nums1 {
		m[num]++
	}

	res := make([]int, 0, n1)
	for _, num := range nums2 {
		if v, ok := m[num]; ok && v > 0 {
			res = append(res, num)
			m[num]--
		}
	}
	return res
}
