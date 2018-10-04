package leetcode0916

func wordSubsets(A []string, B []string) []string {
	m := [26]int{}
	for _, b := range B {
		count := counter(b)
		for i, c := range m {
			if count[i] > c {
				m[i] = count[i]
			}
		}
	}
	res := []string{}
L:
	for _, a := range A {
		count := counter(a)
		for i, c := range m {
			if count[i] > c {
				continue L
			}
		}
		res = append(res, a)
	}
	return res
}

func counter(s string) [26]int {
	count := [26]int{}
	for _, c := range s {
		count[c-'a']++
	}
	return count
}
