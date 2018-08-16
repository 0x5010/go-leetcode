package leetcode0805

func splitArraySameAverage(A []int) bool {
	n := len(A)
	sum, m := 0, n/2
	for _, a := range A {
		sum += a
	}
	possible := false
	for i := 1; i <= m; i++ {
		if sum*i%n == 0 {
			possible = true
			break
		}
	}
	if !possible {
		return false
	}
	sums := make([][]map[int]struct{}, n)
	for i := 0; i < n; i++ {
		sums[i] = make([]map[int]struct{}, m+1)
	}
	for j := 1; j <= m; j++ {
		for i := 0; i < n; i++ {
			sums[i][j] = map[int]struct{}{}
			if j == 1 {
				sums[i][j][A[i]] = struct{}{}
			} else {
				for k := i + 1; k < n; k++ {
					for v := range sums[k][j-1] {
						sums[i][j][v+A[i]] = struct{}{}
					}
				}
			}
		}
	}
	for j := 1; j <= m; j++ {
		if sum*j%n == 0 {
			target := sum * j / n
			for i := 0; i < n; i++ {
				if _, ok := sums[i][j][target]; ok {
					return true
				}
			}
		}
	}
	return false
}
