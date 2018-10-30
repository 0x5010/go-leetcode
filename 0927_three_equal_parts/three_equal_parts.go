package leetcode0927

func threeEqualParts(A []int) []int {
	n := len(A)
	count := 0
	for _, a := range A {
		if a == 1 {
			count++
		}
	}
	if count == 0 {
		return []int{0, n - 1}
	}
	if count%3 != 0 {
		return []int{-1, -1}
	}

	k := count / 3
	p1, p2, p3, count1 := 0, 0, 0, 0
	for i, a := range A {
		if a == 1 {
			if count1 == 0 {
				p1 = i
			} else if count1 == k {
				p2 = i
			} else if count1 == 2*k {
				p3 = i
				break
			}
			count1++
		}
	}
	for p3 < n && A[p1] == A[p2] && A[p2] == A[p3] {
		p1++
		p2++
		p3++
	}
	if p3 == n {
		return []int{p1 - 1, p2}
	}
	return []int{-1, -1}
}
