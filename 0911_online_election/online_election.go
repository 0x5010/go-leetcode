package leetcode0911

import "sort"

type TopVotedCandidate struct {
	m []int
	t []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	n := len(persons)
	m := make([]int, n)
	count := make([]int, n)
	lead := -1
	for i, person := range persons {
		count[person]++
		if i == 0 || count[person] >= count[lead] {
			lead = person
		}
		m[i] = lead
	}
	return TopVotedCandidate{
		m: m,
		t: times,
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	return this.m[sort.Search(len(this.t), func(i int) bool { return this.t[i] > t })-1]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
