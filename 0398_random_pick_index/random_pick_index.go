package leetcode0398

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (this *Solution) Pick(target int) int {
	res := -1
	count := 0
	for i, num := range this.nums {
		if num == target {
			count++
			if rand.Intn(count) == 0 {
				res = i
			}
		}
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
