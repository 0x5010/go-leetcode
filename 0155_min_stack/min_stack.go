package leetcode0155

type MinStack struct {
	s []entry
}

type entry struct {
	value, min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	min := x
	if len(this.s) > 0 {
		tmp := this.GetMin()
		if tmp < min {
			min = tmp
		}
	}
	this.s = append(this.s, entry{x, min})
}

func (this *MinStack) Pop() {
	this.s = this.s[:len(this.s)-1]
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1].value
}

func (this *MinStack) GetMin() int {
	return this.s[len(this.s)-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
