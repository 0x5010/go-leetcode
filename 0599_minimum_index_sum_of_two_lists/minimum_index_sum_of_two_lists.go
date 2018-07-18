package leetcode0599

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	if len(list1) > len(list2) {
		list1, list2 = list2, list1
	}

	m := make(map[string]int, len(list2))
	for i, s := range list2 {
		m[s] = i
	}

	min := math.MaxInt32
	res := []string{}
	for i, s := range list1 {
		if j, ok := m[s]; ok {
			if min == i+j {
				res = append(res, s)
			} else if min > i+j {
				min = i + j
				res = []string{s}
			}
		}
	}
	return res
}
