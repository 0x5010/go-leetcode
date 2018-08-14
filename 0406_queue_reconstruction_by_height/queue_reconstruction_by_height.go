package leetcode0406

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	res := make([][]int, 0, len(people))

	insert := func(idx int, person []int) {
		res = append(res, person)
		if len(res)-1 == idx {
			return
		}
		copy(res[idx+1:], res[idx:])
		res[idx] = person
	}
	for _, p := range people {
		insert(p[1], p)
	}
	return res
}
