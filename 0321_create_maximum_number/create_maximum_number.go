package leetcode0321

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	m, n := len(nums1), len(nums2)

	res := make([]int, k)
	for i := 0; i <= m && i <= k; i++ {
		if n < k-i {
			continue
		}

		tmp := combine(maxK(nums1, i), maxK(nums2, k-i))
		if more(tmp, res) {
			copy(res, tmp)
		}
	}
	return res
}

func maxK(nums []int, k int) []int {
	n := len(nums)
	if k == n {
		return nums
	}

	res := make([]int, k)

	index := 0
	for i := 0; i < k; i++ {
		res[i] = nums[index]
		for j := index + 1; j <= n-k+i; j++ {
			if res[i] < nums[j] {
				res[i] = nums[j]
				index = j
			}
		}
		index++
	}
	return res
}

func combine(nums1, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	res := make([]int, 0, m+n)
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] > nums2[j] || more(nums1[i:], nums2[j:]) {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}

	res = append(res, nums1[i:]...)
	res = append(res, nums2[j:]...)
	return res
}

func more(nums1, nums2 []int) bool {
	m, n := len(nums1), len(nums2)
	i := 0
	for ; i < m && i < n; i++ {
		if nums1[i] != nums2[i] {
			break
		}
	}
	return i == n || (i != m && nums1[i] > nums2[i])
}
