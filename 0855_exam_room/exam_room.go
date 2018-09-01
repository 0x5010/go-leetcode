package leetcode0855

type ExamRoom struct {
	l []int
	n int
}

func Constructor(N int) ExamRoom {
	return ExamRoom{
		l: []int{},
		n: N,
	}
}

func (this *ExamRoom) Seat() int {
	n := len(this.l)
	if n == 0 {
		this.l = append(this.l, 0)
		return 0
	}
	index, d, m := 0, this.l[0], 0
	if this.n-1-this.l[n-1] > d {
		d = this.n - 1 - this.l[n-1]
		index = n
	}
	for i := 1; i < n; i++ {
		cur := (this.l[i] - this.l[i-1]) / 2
		if cur > d || (cur == d && index == n) {
			d = cur
			index = i
			m = (this.l[i-1] + this.l[i]) / 2
		}
	}
	if index == 0 {
		this.l = append([]int{0}, this.l...)
		return 0
	} else if index == n {
		this.l = append(this.l, this.n-1)
		return this.n - 1
	}
	this.l = append(this.l, 0)
	copy(this.l[index+1:], this.l[index:])
	this.l[index] = m
	return m
}

func (this *ExamRoom) Leave(p int) {
	n := len(this.l)
	l, r := 0, n-1
	for l <= r {
		m := (l + r) / 2
		if this.l[m] == p {
			copy(this.l[m:], this.l[m+1:])
			this.l = this.l[:len(this.l)-1]
			return
		} else if this.l[m] < p {
			l = m + 1
		} else {
			r = m - 1
		}
	}
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(N);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
