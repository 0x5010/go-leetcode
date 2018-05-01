package leetcode0376

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	k := 0
	for k < n-1 && nums[k] == nums[k+1] {
		k++
	}
	if k == n-1 {
		return 1
	}
	res := 2
	small := nums[k] < nums[k+1]
	for i := k + 1; i < n-1; i++ {
		if (small && nums[i+1] < nums[i]) ||
			(!small && nums[i+1] > nums[i]) {
			nums[res] = nums[i+1]
			res++
			small = !small
		}
	}
	return res
}
