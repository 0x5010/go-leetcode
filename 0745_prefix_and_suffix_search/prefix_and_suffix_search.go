package leetcode0745

type WordFilter struct {
	m map[string]int
}

func Constructor(words []string) WordFilter {
	m := map[string]int{}
	for k := 0; k < len(words); k++ {
		for i := 0; i <= 10 && i <= len(words[k]); i++ {
			for j := len(words[k]); 0 <= j && len(words[k])-10 <= j; j-- {
				m[words[k][:i]+"#"+words[k][j:]] = k
			}
		}
	}
	return WordFilter{m: m}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	if index, ok := this.m[prefix+"#"+suffix]; ok {
		return index
	}
	return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
