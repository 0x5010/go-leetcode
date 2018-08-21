package leetcode0786

func kthSmallestPrimeFraction(A []int, K int) []int {
	n := len(A)
	l, h, p, q := 0., 1.0, 0, 1
	for l < h {
		mid := (l + h) / 2
		count := 0
		p = 0
		for i, j := 0, n-1; i < n; i++ {
			for j >= 0 && float64(A[i]) > mid*float64(A[n-1-j]) {
				j--
			}
			count += j + 1
			if j >= 0 && p*A[n-1-j] < q*A[i] {
				p, q = A[i], A[n-1-j]
			}
		}
		if count == K {
			return []int{p, q}
		} else if count < K {
			l = mid
		} else {
			h = mid
		}
	}
	return []int{p, q}
}
