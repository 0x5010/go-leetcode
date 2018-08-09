package leetcode0762

import "math/bits"

func countPrimeSetBits(L int, R int) int {
	res := 0
	for i := L; i <= R; i++ {
		switch bits.OnesCount32(uint32(i)) {
		case 2, 3, 5, 7, 11, 13, 17, 19:
			res++
		}
	}
	return res
}
