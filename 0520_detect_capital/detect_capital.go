package leetcode0520

func detectCapitalUse(word string) bool {
	if len(word) < 2 {
		return true
	}
	first, second := isCapital(word[0]), isCapital(word[1])
	if !second {
		return detectOther(word, false)
	} else if first {
		return detectOther(word, true)
	}
	return false
}

func detectOther(word string, capital bool) bool {
	for i := 2; i < len(word); i++ {
		if isCapital(word[i]) != capital {
			return false
		}
	}
	return true
}

func isCapital(b byte) bool {
	return b >= 'A' && b <= 'Z'
}
