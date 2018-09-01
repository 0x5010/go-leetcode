package leetcode0859

func buddyStrings(A string, B string) bool {
	n := len(A)
	if len(B) != n {
		return false
	}
	if A == B {
		s := map[byte]struct{}{}
		for i := 0; i < n; i++ {
			s[A[i]] = struct{}{}
		}
		return len(s) < n
	}
	i, count := 0, 2
	var a, b byte
	for count > 0 && i < n {
		if A[i] != B[i] {
			a += A[i]
			b += B[i]
			count--
		}
		i++
	}
	return a == b && A[i:] == B[i:]
}
