package leetcode0888

func fairCandySwap(A []int, B []int) []int {
	sumA, sumB := 0, 0
	inA := map[int]struct{}{}
	for _, a := range A {
		sumA += a
		inA[a] = struct{}{}
	}
	for _, b := range B {
		sumB += b
	}
	m := (sumA - sumB) / 2
	a, b := 0, 0
	for _, b = range B {
		a = b + m
		if _, ok := inA[a]; ok {
			break
		}
	}
	return []int{a, b}
}
