package leetcode0432

import "container/heap"

type AllOne struct {
	m    map[string]*entry
	maxh maxHeap
	minh minHeap
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{m: make(map[string]*entry),
		maxh: maxHeap{},
		minh: minHeap{},
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	if e, ok := this.m[key]; ok {
		e.v++
		heap.Fix(&this.maxh, e.maxIndex)
		heap.Fix(&this.minh, e.minIndex)
	} else {
		e = &entry{k: key, v: 1}
		this.m[key] = e
		heap.Push(&this.maxh, e)
		heap.Push(&this.minh, e)
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	if e, ok := this.m[key]; !ok {
		return
	} else if e.v == 1 {
		delete(this.m, key)
		heap.Remove(&this.maxh, e.maxIndex)
		heap.Remove(&this.minh, e.minIndex)
	} else {
		e.v--
		heap.Fix(&this.maxh, e.maxIndex)
		heap.Fix(&this.minh, e.minIndex)
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if len(this.maxh) == 0 {
		return ""
	}
	return this.maxh[0].k
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if len(this.minh) == 0 {
		return ""
	}
	return this.minh[0].k
}

type entry struct {
	k        string
	v        int
	maxIndex int
	minIndex int
}

type maxHeap []*entry

func (mh maxHeap) Len() int           { return len(mh) }
func (mh maxHeap) Less(i, j int) bool { return mh[i].v > mh[j].v }
func (mh maxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
	mh[i].maxIndex, mh[j].maxIndex = i, j
}

func (mh *maxHeap) Push(x interface{}) {
	e := x.(*entry)
	e.maxIndex = len(*mh)
	*mh = append(*mh, e)
}

func (mh *maxHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	e := old[n-1]
	*mh = old[:n-1]
	return e
}

type minHeap []*entry

func (mh minHeap) Len() int           { return len(mh) }
func (mh minHeap) Less(i, j int) bool { return mh[i].v < mh[j].v }
func (mh minHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
	mh[i].minIndex, mh[j].minIndex = i, j
}

func (mh *minHeap) Push(x interface{}) {
	e := x.(*entry)
	e.minIndex = len(*mh)
	*mh = append(*mh, e)
}

func (mh *minHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	e := old[n-1]
	*mh = old[:n-1]
	return e
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
