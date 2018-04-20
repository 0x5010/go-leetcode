package leetcode0327

func countRangeSum(nums []int, lower int, upper int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	sums := make([]int, n+1)
	for i, n := range nums {
		sums[i+1] = sums[i] + n
	}
	return countWhilemergeSort(sums, 0, n+1, lower, upper)
}

func countWhilemergeSort(sums []int, start, end, lower, upper int) int {
	if end-start <= 1 {
		return 0
	}

	mid := (start + end) / 2
	count := countWhilemergeSort(sums, start, mid, lower, upper) +
		countWhilemergeSort(sums, mid, end, lower, upper)
	l, r, t := mid, mid, mid

	j := 0
	cache := make([]int, end-start)
	for i := start; i < mid; i++ {
		for l < end && sums[l]-sums[i] < lower {
			l++
		}
		for r < end && sums[r]-sums[i] <= upper {
			r++
		}
		for t < end && sums[t] < sums[i] {
			cache[j] = sums[t]
			j++
			t++
		}
		cache[j] = sums[i]
		count += r - l
		j++
	}
	copy(sums[start:], cache[0:t-start])
	return count
}
