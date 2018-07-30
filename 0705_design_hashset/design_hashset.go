package leetcode0705

type MyHashSet struct {
	s        []bool
	len, cap int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		s:   make([]bool, 10000, 1000001),
		len: 10000,
		cap: 1000001,
	}
}

func (this *MyHashSet) Add(key int) {
	n := this.len
	if n <= key {
		l := 2 * key
		if l > this.cap {
			l = this.cap
		}
		this.len = l
		this.s = this.s[:l]
	}
	this.s[key] = true
}

func (this *MyHashSet) Remove(key int) {
	if key >= this.len {
		return
	}
	this.s[key] = false
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	if key >= this.len {
		return false
	}
	return this.s[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
