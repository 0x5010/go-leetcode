package leetcode0845

func longestMountain(A []int) int {
	n := len(A)
	up, down := 0, 0
	res := 0
	for i := 1; i < n; i++ {
		if down != 0 && A[i-1] < A[i] {
			up, down = 1, 0
		} else if A[i-1] == A[i] {
			up, down = 0, 0
		} else if A[i-1] < A[i] {
			up++
		} else if A[i-1] > A[i] {
			down++
		}
		if up > 0 && down > 0 {
			res = max(res, up+1+down)
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
