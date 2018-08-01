package leetcode0715

import "sort"

type RangeModule struct {
	r []*interval
}

func Constructor() RangeModule {
	return RangeModule{
		r: []*interval{},
	}
}

func (this *RangeModule) AddRange(left int, right int) {
	it := &interval{l: left, r: right}
	n := len(this.r)
	i := sort.Search(n, func(i int) bool {
		return left <= this.r[i].r
	})
	j := i
	for j < n && this.r[j].l <= right {
		it.merge(this.r[j])
		j++
	}

	if i == j {
		this.r = append(this.r, nil)
	}
	copy(this.r[i+1:], this.r[j:])
	this.r = this.r[:n-j+i+1]
	this.r[i] = it
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	n := len(this.r)
	i := sort.Search(n, func(i int) bool {
		return right <= this.r[i].r
	})
	return 0 <= i && i < n && this.r[i].cover(left, right)
}

func (this *RangeModule) RemoveRange(left int, right int) {
	it := &interval{l: left, r: right}
	n := len(this.r)
	if n == 0 {
		return
	}
	i := sort.Search(n, func(i int) bool {
		return left < this.r[i].r
	})
	tmp := []*interval{}
	j := i
	for j < n && this.r[j].l < right {
		ia, ib := minus(this.r[j], it)
		if ia != nil {
			tmp = append(tmp, ia)
		}
		if ib != nil {
			tmp = append(tmp, ib)
		}
		j++
	}
	this.r = append(this.r, nil)
	copy(this.r[i+len(tmp):], this.r[j:])
	this.r = this.r[:n-j+i+len(tmp)]
	for k := 0; k < len(tmp); k++ {
		this.r[i+k] = tmp[k]
	}
}

type interval struct {
	l, r int
}

func (i *interval) cover(l, r int) bool {
	return i.l <= l && r <= i.r
}

func (i *interval) merge(other *interval) {
	if other.l < i.l {
		i.l = other.l
	}
	if i.r < other.r {
		i.r = other.r
	}
}

func minus(a, b *interval) (*interval, *interval) {
	if b.l <= a.l && a.r <= b.r {
		return nil, nil
	}

	if b.l <= a.l && a.l < b.r && b.r < a.r {
		return &interval{l: b.r, r: a.r}, nil
	}

	if a.l < b.l && b.l < a.r && a.r < b.r {
		return &interval{l: a.l, r: b.l}, nil
	}

	return &interval{l: a.l, r: b.l}, &interval{l: b.r, r: a.r}
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
