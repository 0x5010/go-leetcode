package leetcode0873

import "sort"

func lenLongestFibSubseq(A []int) int {
	n := len(A)
	res := 0
	for k := 2; k < n && k < n-res+2; k++ {
		l, r := 0, k-1
		for l < r {
			s := A[l] + A[r]
			if s == A[k] {
				count := 3
				i, j := r, k
				for {
					next := A[i] + A[j]
					i, j = j, j+sort.SearchInts(A[j:], next)
					if j == n || A[j] != next {
						break
					}
					count++
				}
				if count > res {
					res = count
				}
				l++
				r--
			} else if s < A[k] {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
