package leetcode0088

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j, k := m+n-1, m-1, n-1; i >= 0; i-- {
		chosen := 0
		if k < 0 {
			chosen = nums1[j]
			j--
		} else if j < 0 {
			chosen = nums2[k]
			k--
		} else {
			if nums1[j] > nums2[k] {
				chosen = nums1[j]
				j--
			} else {
				chosen = nums2[k]
				k--
			}
		}

		if i < len(nums1) {
			nums1[i] = chosen
		}
	}
}
