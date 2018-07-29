package leetcode0697

func findShortestSubArray(nums []int) int {
	m := map[int]*entry{}
	degree := 0

	res := 0
	for i, num := range nums {
		if e, ok := m[num]; ok {
			e.count++
			e.r = i
			tmp := e.r - e.l + 1
			if e.count > degree {
				degree = e.count
				res = tmp
			} else if e.count == degree && tmp < res {
				res = tmp
			}
		} else {
			m[num] = &entry{
				count: 1,
				l:     i,
				r:     i,
			}
			if res == 0 {
				res = 1
				degree = 1
			}
		}
	}
	return res
}

type entry struct {
	count, l, r int
}
