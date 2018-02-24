package leetcode0078

func subsets(nums []int) [][]int {
	res := [][]int{}

	var recur func([]int, []int)
	recur = func(nums, tmp []int) {
		n := len(nums)
		if n == 0 {
			res = append(res, tmp)
			return
		}
		recur(nums[:n-1], tmp)
		recur(nums[:n-1], append([]int{nums[n-1]}, tmp...))
	}
	recur(nums, []int{})
	return res
}
