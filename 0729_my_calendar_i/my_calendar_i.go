package leetcode0729

type MyCalendar struct {
	head *entry
}

type entry struct {
	start, end int
	next       *entry
}

func Constructor() MyCalendar {
	return MyCalendar{head: &entry{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
	e := &entry{start: start, end: end}
	if this.head.next == nil || e.end <= this.head.next.start {
		e.next, this.head.next = this.head.next, e
		return true
	}

	p := this.head
	for p.next != nil && p.next.end <= e.start {
		p = p.next
	}
	if p.next == nil || e.end <= p.next.start {
		e.next, p.next = p.next, e
		return true
	}
	return false
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
