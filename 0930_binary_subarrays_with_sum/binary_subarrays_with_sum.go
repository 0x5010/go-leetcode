package leetcode0930

func numSubarraysWithSum(A []int, S int) int {
	n := len(A)
	preSum := make([]int, n+1)
	sum := 0
	res := 0
	for _, a := range A {
		sum += a
		tmp := sum - S
		if tmp >= 0 {
			res += preSum[tmp]
		}
		if sum == S {
			res++
		}
		preSum[sum]++
	}
	return res
}
