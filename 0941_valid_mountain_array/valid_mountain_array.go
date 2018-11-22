package leetcode0941

func validMountainArray(A []int) bool {
	n := len(A)
	i, j := 0, n-1
	for i+1 < n && A[i] < A[i+1] {
		i++
	}
	for j > 0 && A[j-1] > A[j] {
		j--
	}
	return i == j && i > 0 && j < n-1
}
