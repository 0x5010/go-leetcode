package leetcode0352

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type SummaryRanges struct {
	is []Interval
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	return SummaryRanges{}
}

func (this *SummaryRanges) Addnum(val int) {
	if this.is == nil {
		this.is = []Interval{
			Interval{
				Start: val,
				End:   val,
			},
		}
		return
	}

	i, j := 0, len(this.is)-1
	for i <= j {
		mid := (i + j) / 2
		if this.is[mid].Start <= val && val <= this.is[mid].End {
			return
		} else if val < this.is[mid].Start {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}

	if i == 0 {
		if this.is[0].Start-1 == val {
			this.is[0].Start--
			return
		}
		ni := Interval{Start: val, End: val}
		this.is = append(this.is, ni)
		copy(this.is[1:], this.is)
		this.is[0] = ni
		return
	}

	if i == len(this.is) {
		if this.is[i-1].End+1 == val {
			this.is[i-1].End++
			return
		}
		this.is = append(this.is, Interval{Start: val, End: val})
		return
	}

	if this.is[i-1].End+1 < val && val < this.is[i].Start-1 {
		this.is = append(this.is, Interval{})
		copy(this.is[i+1:], this.is[i:])
		this.is[i] = Interval{Start: val, End: val}
		return
	}

	if this.is[i-1].End == val-1 && val+1 == this.is[i].Start {
		this.is[i-1].End = this.is[i].End
		n := len(this.is)
		copy(this.is[i:], this.is[i+1:])
		this.is = this.is[:n-1]
		return
	}

	if this.is[i-1].End == val-1 {
		this.is[i-1].End++
	} else {
		this.is[i].Start--
	}
}

func (this *SummaryRanges) Getintervals() []Interval {
	return this.is
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Addnum(val);
 * param_2 := obj.Getintervals();
 */
