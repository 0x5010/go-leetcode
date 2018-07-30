package leetcode0706

type MyHashMap struct {
	m        []int
	len, cap int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		m:   make([]int, 10000, 1000001),
		len: 10000,
		cap: 1000001,
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	if this.len <= key {
		l := 2 * key
		if l > this.cap {
			l = this.cap
		}
		this.len = l
		this.m = this.m[:l]
	}
	this.m[key] = value + 1
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	if key >= this.len {
		return -1
	}
	return this.m[key] - 1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	if key >= this.len {
		return
	}
	this.m[key] = 0
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
