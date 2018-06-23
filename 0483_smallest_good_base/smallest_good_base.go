package leetcode0483

import (
	"math"
	"strconv"
)

func smallestGoodBase(n string) string {
	num, _ := strconv.ParseUint(n, 10, 64)
	m := int(math.Log2(float64(num)))
	for i := m; i >= 1; i-- {
		k := uint64(math.Pow(float64(num), 1.0/float64(i)))
		var sum, a uint64 = 1, 1
		for j := 1; j <= i; j++ {
			a *= k
			sum += a
		}
		if sum == num {
			return strconv.FormatUint(k, 10)
		}
	}
	return strconv.FormatUint(num-1, 10)
}
