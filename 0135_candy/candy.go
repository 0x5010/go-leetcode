package leetcode0135

func candy(ratings []int) int {
	n := len(ratings)
	if n < 2 {
		return n
	}
	num := make([]int, n)
	for i := 0; i < n; i++ {
		num[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			num[i] = num[i-1] + 1
		}
	}
	for i := n - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && num[i]+1 > num[i-1] {
			num[i-1] = num[i] + 1
		}
	}

	sum := 0
	for _, i := range num {
		sum += i
	}
	return sum
}
