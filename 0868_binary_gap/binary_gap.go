package leetcode0868

func binaryGap(N int) int {
	gap := 0
	res := 0
	for N > 0 {
		if N&1 == 1 {
			if gap > res {
				res = gap
			}
			gap = 1
		} else if gap > 0 {
			gap++
		}
		N >>= 1
	}
	return res
}
