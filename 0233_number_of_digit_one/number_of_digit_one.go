package leetcode0233

func countDigitOne(n int) int {
	res := 0
	m := 1
	for m <= n {
		res += (n/m + 8) / 10 * m
		if n/m%10 == 1 {
			res += n%m + 1
		}
		m *= 10
	}
	return res
}
