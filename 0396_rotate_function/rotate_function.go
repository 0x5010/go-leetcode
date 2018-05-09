package leetcode0396

func maxRotateFunction(A []int) int {
	n := len(A)
	f, sum := 0, 0
	for i, a := range A {
		f += i * a
		sum += a
	}

	res := f
	for i := n - 1; i >= 1; i-- {
		f += sum - n*A[i]
		if f > res {
			res = f
		}
	}
	return res
}
