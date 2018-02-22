package leetcode0068

import (
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	res := []string{}

	for i, w := 0, 0; i < len(words); i = w {
		l := -1
		for w = i; w < len(words) && l+len(words[w])+1 <= maxWidth; w++ {
			l += len(words[w]) + 1
		}

		str := words[i]
		space, extra := 1, 0
		if w != i+1 && w != len(words) {
			space = (maxWidth-l)/(w-i-1) + 1
			extra = (maxWidth - l) % (w - i - 1)
		}
		for j := i + 1; j < w; j++ {
			str += strings.Repeat(" ", space)
			if extra > 0 {
				str += " "
				extra--
			}
			str += words[j]
		}
		str += strings.Repeat(" ", maxWidth-len(str))
		res = append(res, str)
	}
	return res
}
