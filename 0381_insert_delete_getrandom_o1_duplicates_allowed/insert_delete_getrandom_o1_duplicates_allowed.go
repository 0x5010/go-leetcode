package leetcode0381

import (
	"math/rand"
	"sort"
)

type RandomizedCollection struct {
	l []int
	m map[int][]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		l: []int{},
		m: make(map[int][]int),
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	_, ok := this.m[val]
	this.l = append(this.l, val)
	this.m[val] = append(this.m[val], len(this.l)-1)
	return !ok
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	indexs := this.m[val]
	if len(indexs) == 0 {
		return false
	}
	tmp := this.l[len(this.l)-1]
	valIndex := indexs[len(indexs)-1]

	this.l[valIndex] = tmp
	this.m[tmp][len(this.m[tmp])-1] = valIndex
	sort.Ints(this.m[tmp])
	this.l = this.l[:len(this.l)-1]
	this.m[val] = indexs[:len(indexs)-1]

	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	return this.l[rand.Intn(len(this.l))]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
