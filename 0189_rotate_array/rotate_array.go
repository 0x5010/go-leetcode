package leetcode0189

func rotate(nums []int, k int) {
	n := len(nums)
	tmp := make([]int, n)
	copy(tmp, nums)
	for i := 0; i < n; i++ {
		nums[(i+k)%n] = tmp[i]
	}
}
