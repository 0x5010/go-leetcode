package leetcode0854

func kSimilarity(A string, B string) int {
	if A == B {
		return 0
	}
	n := len(A)
	seen := map[string]struct{}{}
	seen[A] = struct{}{}
	queue := []string{A}
	res := 0
	for {
		res++
		count := len(queue)
		for k := 0; k < count; k++ {
			s := queue[k]
			i := 0
			for s[i] == B[i] {
				i++
			}
			for j := i + 1; j < n; j++ {
				if s[j] == B[j] || s[i] != B[j] {
					continue
				}
				tmp := swap(s, i, j)
				if tmp == B {
					return res
				}
				if _, ok := seen[tmp]; !ok {
					seen[tmp] = struct{}{}
					queue = append(queue, tmp)
				}
			}
		}
		queue = queue[count:]
	}
}

func swap(s string, i, j int) string {
	bs := []byte(s)
	bs[i], bs[j] = bs[j], bs[i]
	return string(bs)
}
