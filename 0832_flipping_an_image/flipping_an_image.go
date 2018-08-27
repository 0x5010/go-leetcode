package leetcode0832

func flipAndInvertImage(A [][]int) [][]int {
	n := len(A)
	for _, r := range A {
		i, j := 0, n-1
		for i < j {
			r[i], r[j] = r[j]^1, r[i]^1
			i++
			j--
		}
		if i == j {
			r[i] ^= 1
		}
	}
	return A
}
