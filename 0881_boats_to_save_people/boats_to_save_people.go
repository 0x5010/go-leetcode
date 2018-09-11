package leetcode0881

import "sort"

func numRescueBoats(people []int, limit int) int {
	n := len(people)
	sort.Ints(people)

	res := 0
	for l, r := 0, n-1; l <= r; r-- {
		if people[l]+people[r] <= limit {
			l++
		}
		res++
	}
	return res
}
