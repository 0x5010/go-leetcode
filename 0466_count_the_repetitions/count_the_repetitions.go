package leetcode0466

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	count := 0
	s := s1
	// remainString := []byte{}
	m := map[string]int{"": 0}
	countList := []int{0}
	foundLoop := false
	foundIndex := -1
	for i := 0; i <= n1; i++ {
		remainString := []byte{}
		lastMatch, p2 := -1, 0
		for p1 := 0; p1 < len(s); p1++ {
			if s[p1] == s2[p2] {
				p2++
				if p2 == len(s2) {
					p2 = 0
					count++
					lastMatch = p1
				}
			}
		}
		for j := lastMatch + 1; j < len(s); j++ {
			remainString = append(remainString, s[j])
		}
		remain := string(remainString)
		if index, ok := m[remain]; ok {
			foundLoop = true
			foundIndex = index
			break
		}
		m[remain] = i + 1
		countList = append(countList, count)
		s = remain + s1
	}
	if !foundLoop {
		return count / n2
	}
	countOfLoop := count - countList[foundIndex]
	loopLen := len(m) - foundIndex
	count = countList[foundIndex]
	n1 -= foundIndex
	count += countOfLoop * (n1 / loopLen)
	n1 %= loopLen
	count += countList[foundIndex+n1] - countList[foundIndex]
	return count / n2
}
