package leetcode0933

type RecentCounter struct {
	q []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		q: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	for len(this.q) != 0 && this.q[0] < t-3000 {
		this.q = this.q[1:]
	}
	this.q = append(this.q, t)
	return len(this.q)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
