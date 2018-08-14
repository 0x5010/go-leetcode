package leetcode0451

import "sort"

func frequencySort(s string) string {
	n := len(s)
	freqs := make([]freq, 128)
	for i := 0; i < n; i++ {
		b := s[i]
		f := freqs[b]
		f.b = b
		f.f++
		freqs[b] = f
	}
	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i].f > freqs[j].f
	})
	res := make([]byte, n)
	i := 0
	for _, f := range freqs {
		if f.f == 0 {
			break
		}
		for j := 0; j < f.f; j++ {
			res[i] = f.b
			i++
		}
	}
	return string(res)
}

type freq struct {
	b byte
	f int
}
