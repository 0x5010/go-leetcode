package leetcode0848

func shiftingLetters(S string, shifts []int) string {
	bs := []byte(S)
	shift := 0
	for i := len(bs) - 1; i >= 0; i-- {
		shift += shifts[i]
		bs[i] = byte((int(S[i]-'a')+shift)%26 + 'a')
	}
	return string(bs)
}
