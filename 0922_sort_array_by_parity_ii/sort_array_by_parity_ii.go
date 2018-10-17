package leetcode0922

func sortArrayByParityII(A []int) []int {
	n := len(A)
	for i, j := 0, 1; i < n; i += 2 {
		for j < n && A[j]%2 == 1 {
			j += 2
		}
		if A[i]%2 == 1 {
			A[i], A[j] = A[j], A[i]
			j += 2
		}
	}
	return A
}
