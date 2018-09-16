package leetcode0901

type StockSpanner struct {
	s [][2]int
	d int
}

func Constructor() StockSpanner {
	return StockSpanner{
		s: [][2]int{},
	}
}

func (this *StockSpanner) Next(price int) int {
	this.d++
	i := len(this.s)
	for i > 0 && this.s[i-1][0] <= price {
		i--
	}
	this.s = append(this.s[:i], [2]int{price, this.d})
	if i == 0 {
		return this.d
	}
	return this.d - this.s[i-1][1]
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
