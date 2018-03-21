package leetcode0204

import (
	"math"
)

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}

	notPrime := make([]bool, n)
	count := n / 2
	for i := 3; i < int(math.Sqrt(float64(n)))+1; i += 2 {
		if notPrime[i] {
			continue
		}

		for j := i * i; j < n; j += 2 * i {
			if !notPrime[j] {
				count--
				notPrime[j] = true
			}
		}
	}
	return count
}
