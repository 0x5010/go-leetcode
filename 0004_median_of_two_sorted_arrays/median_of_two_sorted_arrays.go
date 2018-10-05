package leetcode0004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := merge(nums1, nums2)
	n := len(nums)
	m := n / 2
	if n%2 == 0 {
		return float64(nums[m-1]+nums[m]) / 2
	}
	return float64(nums[m])
}

func merge(nums1, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	i, j := 0, 0
	res := make([]int, m+n)
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			res[i+j] = nums1[i]
			i++
		} else {
			res[i+j] = nums2[j]
			j++
		}
	}
	if i < m {
		copy(res[i+j:], nums1[i:])
	} else if j < n {
		copy(res[i+j:], nums2[j:])
	}
	return res
}
