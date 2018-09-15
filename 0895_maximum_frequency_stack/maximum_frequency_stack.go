package leetcode0895

type FreqStack struct {
	freq map[int]int
	s    [][]int
}

func Constructor() FreqStack {
	return FreqStack{
		freq: map[int]int{},
	}
}

func (this *FreqStack) Push(x int) {
	i := this.freq[x]
	this.freq[x]++
	if i == len(this.s) {
		this.s = append(this.s, []int{x})
	} else {
		this.s[i] = append(this.s[i], x)
	}
}

func (this *FreqStack) Pop() int {
	i := len(this.s) - 1
	if len(this.s[i]) == 0 {
		this.s = this.s[:i]
		i--
	}
	j := len(this.s[i]) - 1
	x := this.s[i][j]
	this.s[i] = this.s[i][:j]
	this.freq[x]--
	return x
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 */
