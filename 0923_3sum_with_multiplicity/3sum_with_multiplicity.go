package leetcode0923

func threeSumMulti(A []int, target int) int {
	mod := 1000000007
	m := [101]int{}
	for _, a := range A {
		m[a]++
	}
	res := 0
	for i := 0; i < 101; i++ {
		for j := i; j < 101; j++ {
			k := target - i - j
			if k > 100 || k < 0 {
				continue
			}
			if i == j && j == k {
				res += m[i] * (m[i] - 1) * (m[i] - 2) / 6
			} else if i == j {
				res += m[i] * (m[i] - 1) / 2 * m[k]
			} else if j < k {
				res += m[i] * m[j] * m[k]
			}
		}
	}
	return res % mod
}
