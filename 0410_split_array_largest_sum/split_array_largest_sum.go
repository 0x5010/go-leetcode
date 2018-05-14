package leetcode0410

func splitArray(nums []int, m int) int {
	sum, biggest := 0, 0
	for _, n := range nums {
		sum += n
		biggest = max(biggest, n)
	}
	if m == 1 {
		return sum
	}

	bigger := func(guess int) bool {
		count, subsum := 1, 0
		for _, n := range nums {
			subsum += n
			if subsum > guess {
				subsum = n
				count++
				if count > m {
					return false
				}
			}
		}
		return true
	}

	l, r := biggest, sum
	for l <= r {
		mid := (l + r) / 2
		if bigger(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
