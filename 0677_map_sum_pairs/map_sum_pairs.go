package leetcode0677

import (
	"sort"
	"strings"
)

type MapSum struct {
	m    map[string]int
	keys []string
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		m:    map[string]int{},
		keys: []string{},
	}
}

func (this *MapSum) Insert(key string, val int) {
	this.m[key] = val
	i := sort.SearchStrings(this.keys, key)
	if i < len(this.keys) && this.keys[i] == key {
		return
	}
	this.keys = append(this.keys, "")
	copy(this.keys[i+1:], this.keys[i:])
	this.keys[i] = key
}

func (this *MapSum) Sum(prefix string) int {
	sum := 0
	i := sort.SearchStrings(this.keys, prefix)
	for i < len(this.keys) && strings.HasPrefix(this.keys[i], prefix) {
		sum += this.m[this.keys[i]]
		i++
	}
	return sum
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
