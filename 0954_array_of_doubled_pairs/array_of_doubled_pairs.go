package leetcode0954

import "sort"

func canReorderDoubled(A []int) bool {
	c := map[int]int{}
	for _, a := range A {
		c[a]++
	}
	keys := []int{}
	for k := range c {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return abs(keys[i]) < abs(keys[j])
	})
	for _, x := range keys {
		if c[x] > c[2*x] {
			return false
		}
		c[2*x] -= c[x]
	}
	return true
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
