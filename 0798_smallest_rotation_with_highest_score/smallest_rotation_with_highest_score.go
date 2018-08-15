package leetcode0798

func bestRotation(A []int) int {
	n := len(A)
	delta := make([]int, n)
	for i, a := range A {
		delta[(i-a+1+n)%n]--
	}
	res := 0
	for i := 1; i < n; i++ {
		delta[i] += delta[i-1] + 1
		if delta[res] < delta[i] {
			res = i
		}
	}
	return res
}
