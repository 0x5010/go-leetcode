package leetcode0731

type MyCalendarTwo struct {
	events    [][2]int
	conflicts [][2]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		events:    [][2]int{},
		conflicts: [][2]int{},
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, conflict := range this.conflicts {
		if conflict[1] > start && end > conflict[0] {
			return false
		}
	}
	for _, event := range this.events {
		if event[1] > start && end > event[0] {
			this.conflicts = append(this.conflicts, [2]int{
				max(event[0], start),
				min(event[1], end),
			})
		}
	}
	this.events = append(this.events, [2]int{start, end})
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
