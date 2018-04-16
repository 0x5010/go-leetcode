package leetcode0316

func removeDuplicateLetters(s string) string {
	n := len(s)
	stack := []byte{}

	count := [26]int{}
	for i := 0; i < n; i++ {
		count[s[i]-'a']++
	}

	visited := [26]bool{}
	for i := 0; i < n; i++ {
		b := s[i]
		count[b-'a']--
		if visited[b-'a'] {
			continue
		}

		for len(stack) != 0 {
			p := stack[len(stack)-1]
			if p > b && count[p-'a'] > 0 {
				visited[p-'a'] = false
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, b)
		visited[b-'a'] = true
	}

	return string(stack)
}

func reverse(bs []byte) string {
	i, j := 0, len(bs)-1
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return string(bs)
}
