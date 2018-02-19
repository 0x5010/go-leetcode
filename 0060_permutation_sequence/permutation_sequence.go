package leetcode0060

func getPermutation(n int, k int) string {
	if n == 0 {
		return ""
	}

	res := make([]byte, n)

	nums := make([]byte, n)
	for i := 0; i < n; i++ {
		nums[i] = byte(i) + '1'
	}

	f := 1
	for i := 1; i < n; i++ {
		f *= i
	}

	k--

	for i := 0; i < n; i++ {
		index := k / f
		res[i] = nums[index]
		if i == n-1 {
			break
		}
		nums = append(nums[:index], nums[index+1:]...)
		k %= f
		f /= n - i - 1
	}
	return string(res)
}
