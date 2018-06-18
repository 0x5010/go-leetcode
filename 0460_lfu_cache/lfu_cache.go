package leetcode0460

import (
	"container/heap"
	"time"
)

type LFUCache struct {
	m   map[int]*entry
	fh  freqHeap
	cap int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		m:   make(map[int]*entry, capacity),
		fh:  make(freqHeap, 0, capacity),
		cap: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if e, ok := this.m[key]; ok {
		this.fh.update(e)
		return e.v
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap <= 0 {
		return
	}
	if e, ok := this.m[key]; ok {
		this.m[key].v = value
		this.fh.update(e)
		return
	}
	e := &entry{k: key, v: value}
	if len(this.fh) == this.cap {
		tmp := heap.Pop(&this.fh).(*entry)
		delete(this.m, tmp.k)
	}
	this.m[key] = e
	heap.Push(&this.fh, e)
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type entry struct {
	k, v        int
	freq, index int
	time        time.Time
}

type freqHeap []*entry

func (fh freqHeap) Len() int { return len(fh) }

func (fh freqHeap) Less(i, j int) bool {
	if fh[i].freq == fh[j].freq {
		return fh[i].time.Before(fh[j].time)
	}
	return fh[i].freq < fh[j].freq
}

func (fh freqHeap) Swap(i, j int) {
	fh[i], fh[j] = fh[j], fh[i]
	fh[i].index = i
	fh[j].index = j
}

func (fh *freqHeap) Push(x interface{}) {
	n := len(*fh)
	e := x.(*entry)
	e.index = n
	e.time = time.Now()
	*fh = append(*fh, e)
}

func (fh *freqHeap) Pop() interface{} {
	old := *fh
	n := len(old)
	x := old[n-1]
	*fh = old[0 : n-1]
	return x
}

func (fh *freqHeap) update(e *entry) {
	e.freq++
	e.time = time.Now()
	heap.Fix(fh, e.index)
}
