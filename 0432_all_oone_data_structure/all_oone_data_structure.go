package leetcode0432

import (
	"container/list"
)

type AllOne struct {
	m map[string]*list.Element
	l *list.List
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{
		m: make(map[string]*list.Element),
		l: list.New(),
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	e, ok := this.m[key]
	if !ok {
		this.m[key] = this.l.PushFront(&entry{
			k: key,
			v: 1,
		})
		return
	}

	ele := e.Value.(*entry)
	ele.v++
	next := e
	for next != nil && next.Value.(*entry).v <= ele.v {
		next = next.Next()
	}
	if next == nil {
		this.l.MoveToBack(e)
	} else {
		this.l.MoveBefore(e, next)
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	e, ok := this.m[key]
	if !ok {
		return
	}

	ele := e.Value.(*entry)
	if ele.v == 1 {
		delete(this.m, key)
		this.l.Remove(e)
		return
	}
	ele.v--
	prev := e
	for prev != nil && prev.Value.(*entry).v >= ele.v {
		prev = prev.Prev()
	}
	if prev == nil {
		this.l.MoveToFront(e)
	} else {
		this.l.MoveAfter(e, prev)
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	tail := this.l.Back()
	if tail != nil {
		return tail.Value.(*entry).k
	}
	return ""
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	head := this.l.Front()
	if head != nil {
		return head.Value.(*entry).k
	}
	return ""
}

type entry struct {
	k string
	v int
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
