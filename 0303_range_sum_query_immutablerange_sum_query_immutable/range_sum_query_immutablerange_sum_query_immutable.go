package leetcode0303

type NumArray struct {
	dp []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	dp := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		dp[i] = dp[i-1] + nums[i-1]
	}
	return NumArray{dp}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.dp[j+1] - this.dp[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
