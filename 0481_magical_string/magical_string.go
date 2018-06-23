package leetcode0481

func magicalString(n int) int {
	if n <= 0 {
		return 0
	} else if n <= 3 {
		return 1
	}
	l := make([]int, n+1)
	l[0], l[1], l[2] = 1, 2, 2
	res := 1
	i, j, num := 2, 3, 1
	for j < n {
		for k := 0; k < l[i]; k++ {
			l[j] = num
			if num == 1 && j < n {
				res++
			}
			j++
		}
		num ^= 3
		i++
	}
	return res
}
