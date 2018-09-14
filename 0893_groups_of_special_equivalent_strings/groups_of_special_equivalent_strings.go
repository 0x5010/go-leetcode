package leetcode0893

func numSpecialEquivGroups(A []string) int {
	n := len(A[0])
	m := map[[26]int]struct{}{}
	for _, a := range A {
		count := [26]int{}
		i := 0
		for ; i+1 < n; i += 2 {
			count[a[i]-'a']++
			count[a[i+1]-'a'] += 21
		}
		if i < n {
			count[a[i]-'a']++
		}
		m[count] = struct{}{}
	}
	return len(m)
}
