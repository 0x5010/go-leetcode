package leetcode0524

func findLongestWord(s string, d []string) string {
	n := len(s)
	res := ""
	for _, word := range d {
		j, l := 0, len(word)
		for i := 0; i < n; i++ {
			if j < l && s[i] == word[j] {
				j++
			}
		}
		if j == l && l >= len(res) &&
			(l > len(res) || word < res) {
			res = word
		}
	}
	return res
}
