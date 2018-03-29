package leetcode0238

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	res[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		res[i] = res[i+1] * nums[i+1]
	}
	tmp := 1
	for i, num := range nums {
		res[i] *= tmp
		tmp *= num
	}
	return res
}
