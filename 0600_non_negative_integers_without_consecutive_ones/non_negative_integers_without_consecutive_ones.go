package leetcode0600

func findIntegers(num int) int {
	br := []int{}
	for num > 0 {
		br = append(br, num%2)
		num /= 2
	}
	n := len(br)
	a, b := make([]int, n), make([]int, n)
	a[0], b[0] = 1, 1

	for i := 1; i < n; i++ {
		a[i] = a[i-1] + b[i-1]
		b[i] = a[i-1]
	}

	res := a[n-1] + b[n-1]
	for i := n - 2; i >= 0; i-- {
		if br[i] == 1 && br[i+1] == 1 {
			break
		}
		if br[i] == 0 && br[i+1] == 0 {
			res -= b[i]
		}
	}
	return res
}
