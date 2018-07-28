package leetcode0689

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums) - k + 1

	sum := make([]int, n)
	sumK := 0
	for i := 0; i < len(nums); i++ {
		sumK += nums[i]
		if i >= k {
			sumK -= nums[i-k]
		}
		if i >= k-1 {
			sum[i-k+1] = sumK
		}
	}

	l := make([]int, n)
	indexOfMax := 0
	for i := 0; i < n; i++ {
		if sum[indexOfMax] < sum[i] {
			indexOfMax = i
		}
		l[i] = indexOfMax
	}

	indexOfMax = n - 1
	r := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if sum[indexOfMax] < sum[i] {
			indexOfMax = i
		}
		r[i] = indexOfMax
	}

	a, b, c := 0, k, k+k
	for y := k; y < n-k; y++ {
		x, z := l[y-k], r[y+k]
		if sum[a]+sum[b]+sum[c] < sum[x]+sum[y]+sum[z] {
			a, b, c = x, y, z
		}
	}
	return []int{a, b, c}
}
