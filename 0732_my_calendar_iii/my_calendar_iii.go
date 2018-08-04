package leetcode0732

type MyCalendarThree struct {
	events [][2]int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		events: [][2]int{},
	}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	index := 0
	for i := 0; i < len(this.events); i++ {
		if this.events[i][0] == start {
			this.events[i][1]++
			index = -1
			break
		} else if this.events[i][0] > start {
			break
		}
		index = i + 1
	}
	if index != -1 {
		this.events = append(this.events, [2]int{})
		for i := len(this.events) - 1; i >= index+1; i-- {
			this.events[i] = this.events[i-1]
		}
		this.events[index] = [2]int{start, 1}
	}

	index = 0
	for i := 0; i < len(this.events); i++ {
		if this.events[i][0] == end {
			this.events[i][1]--
			index = -1
			break
		} else if this.events[i][0] > end {
			break
		}
		index = i + 1
	}
	if index != -1 {
		this.events = append(this.events, [2]int{})
		for i := len(this.events) - 1; i >= index+1; i-- {
			this.events[i] = this.events[i-1]
		}
		this.events[index] = [2]int{end, -1}
	}

	count := 0
	res := 0
	for i := 0; i < len(this.events); i++ {
		count += this.events[i][1]
		if count > res {
			res = count
		}
	}
	return res
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
