package leetcode0467

func findSubstringInWraproundString(p string) int {
	counts := [26]int{}
	length := 0
	for i := 0; i < len(p); i++ {
		if 0 < i && (p[i-1]+1 == p[i] || p[i-1] == p[i]+25) {
			length++
		} else {
			length = 1
		}

		b := p[i] - 'a'
		if length > counts[b] {
			counts[b] = length
		}
	}

	res := 0
	for i := 0; i < 26; i++ {
		res += counts[i]
	}
	return res
}
