package leetcode0775

func isIdealPermutation(A []int) bool {
	for i := 0; i < len(A)-1; i++ {
		if A[i] == i {
			continue
		}
		if A[i] == i+1 && A[i+1] == i {
			i++
		} else {
			return false
		}
	}
	return true
}
