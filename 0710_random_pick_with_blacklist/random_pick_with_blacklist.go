package leetcode0710

import "math/rand"

type Solution struct {
	M int
	m map[int]int
}

func Constructor(N int, blacklist []int) Solution {
	M := N - len(blacklist)
	m := make(map[int]int, len(blacklist))
	for _, b := range blacklist {
		m[b] = b
	}
	for _, b := range blacklist {
		if b >= M {
			continue
		}

		N--
		for m[N] == N {
			N--
		}
		m[b] = N
	}

	return Solution{
		M: M,
		m: m,
	}

}

func (this *Solution) Pick() int {
	r := rand.Intn(this.M)
	if t, ok := this.m[r]; ok {
		return t
	}
	return r
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(N, blacklist);
 * param_1 := obj.Pick();
 */
