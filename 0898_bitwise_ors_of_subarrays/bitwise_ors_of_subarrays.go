package leetcode0898

func subarrayBitwiseORs(A []int) int {
	m := map[int]struct{}{}
	s1, s2 := []int{}, []int{}
	for _, a := range A {
		s := map[int]struct{}{a: struct{}{}}
		m[a] = struct{}{}
		s2 = append(s2, a)
		for _, b := range s1 {
			b |= a
			if _, ok := s[b]; !ok {
				s[b] = struct{}{}
				m[b] = struct{}{}
				s2 = append(s2, b)
			}
		}
		s1, s2 = s2, s1[:0]
	}
	return len(m)
}
