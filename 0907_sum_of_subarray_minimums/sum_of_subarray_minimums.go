package leetcode0907

func sumSubarrayMins(A []int) int {
	mod := 1000000007
	n := len(A)
	l, r := make([]int, n), make([]int, n)
	s1, s2 := [][2]int{}, [][2]int{}
	for i := 0; i < n; i++ {
		count := 1
		for len(s1) != 0 && s1[len(s1)-1][0] > A[i] {
			count += s1[len(s1)-1][1]
			s1 = s1[:len(s1)-1]
		}
		s1 = append(s1, [2]int{A[i], count})
		l[i] = count
	}
	for i := n - 1; i >= 0; i-- {
		count := 1
		for len(s2) != 0 && s2[len(s2)-1][0] >= A[i] {
			count += s2[len(s2)-1][1]
			s2 = s2[:len(s2)-1]
		}
		s2 = append(s2, [2]int{A[i], count})
		r[i] = count
	}
	res := 0
	for i := 0; i < n; i++ {
		res = (res + A[i]*l[i]*r[i]) % mod
	}
	return res
}
