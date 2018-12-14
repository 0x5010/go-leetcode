package leetcode0949

import "fmt"

func largestTimeFromDigits(A []int) string {
	res := ""
	var permute func(int, int)
	permute = func(l, r int) {
		if l == r {
			if A[0]*10+A[1] < 24 && A[2] < 6 {
				tmp := fmt.Sprintf("%d%d:%d%d", A[0], A[1], A[2], A[3])
				if tmp > res {
					res = tmp
				}
			}
		}
		for i := l; i <= r; i++ {
			A[i], A[l] = A[l], A[i]
			permute(l+1, r)
			A[i], A[l] = A[l], A[i]
		}
	}
	permute(0, len(A)-1)
	return res
}
