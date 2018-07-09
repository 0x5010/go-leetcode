package leetcode0532

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}

	res := 0
	if k == 0 {
		for _, count := range m {
			if count > 1 {
				res++
			}
		}
	} else {
		for num := range m {
			if m[num-k] > 0 {
				res++
			}
		}
	}

	return res
}
