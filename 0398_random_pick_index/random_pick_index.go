package leetcode0398

import "math/rand"

type Solution struct {
	m map[int][]int
}

func Constructor(nums []int) Solution {
	m := make(map[int][]int)
	for i, num := range nums {
		m[num] = append(m[num], i)
	}
	return Solution{m: m}
}

func (this *Solution) Pick(target int) int {
	return this.m[target][rand.Intn(len(this.m[target]))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
