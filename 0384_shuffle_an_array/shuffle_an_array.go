package leetcode0384

import "math/rand"

type Solution struct {
	o []int
}

func Constructor(nums []int) Solution {
	return Solution{
		o: nums,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.o
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	n := make([]int, len(this.o))
	copy(n, this.o)
	for i := range n {
		j := rand.Intn(i + 1)
		n[i], n[j] = n[j], n[i]
	}
	return n
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
