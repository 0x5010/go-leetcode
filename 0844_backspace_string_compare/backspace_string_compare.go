package leetcode0844

func backspaceCompare(S string, T string) bool {
	var s, t byte
	for i, j := len(S)-1, len(T)-1; i >= 0 || j >= 0; {
		i, s = next(i, S)
		j, t = next(j, T)
		if s != t {
			return false
		}
	}
	return true
}

func next(index int, s string) (int, byte) {
	for b := 0; index >= 0; index-- {
		if s[index] == '#' {
			b++
		} else if b > 0 {
			b--
		} else {
			return index - 1, s[index]
		}
	}
	return -1, 0
}
