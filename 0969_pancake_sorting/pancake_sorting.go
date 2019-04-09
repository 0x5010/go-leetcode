package leetcode0969

func pancakeSort(A []int) []int {
	n := len(A)
	res := make([]int, 0, n)
	for x, i := n, 0; x > 0; x-- {
		for i = 0; A[i] != x; i++ {
		}
		reverse(A, i+1)
		res = append(res, i+1)
		reverse(A, x)
		res = append(res, x)
	}
	return res
}

func reverse(A []int, k int) {
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		A[i], A[j] = A[j], A[i]
	}
}
