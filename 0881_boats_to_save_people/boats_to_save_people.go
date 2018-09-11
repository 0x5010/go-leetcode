package leetcode0881

import "sort"

func numRescueBoats(people []int, limit int) int {
	n := len(people)
	sort.Ints(people)
	l, r := 0, n-1
	res := 0
	for l <= r {
		if people[l]+people[r] <= limit {
			l++
		}
		r--
		res++
	}
	return res
}
