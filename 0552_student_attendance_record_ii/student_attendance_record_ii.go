package leetcode0552

func checkRecord(n int) int {
	m := 1000000007
	switch n {
	case 0:
		return 1
	case 1:
		return 3
	case 2:
		return 8
	}

	P := make([]int, n)
	P[0] = 1
	L := make([]int, n)
	L[0], L[1] = 1, 3
	A := make([]int, n)
	A[0], A[1], A[2] = 1, 2, 4

	for i := 1; i < n; i++ {
		P[i-1] %= m
		L[i-1] %= m
		A[i-1] %= m

		P[i] = P[i-1] + L[i-1] + A[i-1]

		if i > 1 {
			L[i] = P[i-1] + P[i-2] + A[i-2] + A[i-1]
		}

		if i > 2 {
			A[i] = A[i-1] + A[i-2] + A[i-3]
		}
	}

	return (P[n-1] + L[n-1] + A[n-1]) % m
}
