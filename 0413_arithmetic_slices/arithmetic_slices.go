package leetcode0413

func numberOfArithmeticSlices(A []int) int {
	n := len(A)
	if n < 3 {
		return 0
	}
	res := 0

	for i, j := 0, 0; i < n; {
		j = i + 2
		for j < n && A[j]-A[j-1] == A[j-1]-A[j-2] {
			j++
		}
		j--
		res += (j - i - 1) * (j - i) / 2
		i = j
	}
	return res
}
