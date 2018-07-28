package leetcode0692

import "sort"

func topKFrequent(words []string, k int) []string {
	f := map[string]int{}
	for _, word := range words {
		f[word]++
	}
	l := make([]*entry, 0, len(f))
	for word, freq := range f {
		l = append(l, &entry{
			word: word,
			freq: freq,
		})
	}
	sort.Slice(l, func(i, j int) bool {
		if l[i].freq == l[j].freq {
			return l[i].word < l[j].word
		}
		return l[i].freq > l[j].freq
	})
	res := make([]string, k)
	for i := 0; i < k; i++ {
		res[i] = l[i].word
	}
	return res
}

type entry struct {
	word string
	freq int
}
