package leetcode0905

func sortArrayByParity(A []int) []int {
	n := len(A)
	res := make([]int, n)
	l, r := 0, n-1
	for i := 0; i < n; i++ {
		if A[i]%2 == 0 {
			res[l] = A[i]
			l++
		} else {
			res[r] = A[i]
			r--
		}
	}
	return res
}
