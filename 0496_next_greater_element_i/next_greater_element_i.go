package leetcode0496

func nextGreaterElement(findNums []int, nums []int) []int {
	m := make(map[int]int)
	s := []int{}
	for _, num := range nums {
		for len(s) != 0 && s[len(s)-1] < num {
			m[s[len(s)-1]] = num
			s = s[:len(s)-1]
		}
		s = append(s, num)
	}

	res := make([]int, len(findNums))
	for i, num := range findNums {
		if v, ok := m[num]; ok {
			res[i] = v
		} else {
			res[i] = -1
		}
	}
	return res
}
