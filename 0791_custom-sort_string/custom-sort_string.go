package leetcode0791

func customSortString(S string, T string) string {
	count := [26]int{}
	for i := 0; i < len(T); i++ {
		count[T[i]-'a']++
	}
	bs := []byte{}
	for i := 0; i < len(S); i++ {
		for count[S[i]-'a'] > 0 {
			bs = append(bs, S[i])
			count[S[i]-'a']--
		}
	}
	for i, c := range count {
		for c > 0 {
			bs = append(bs, byte(i)+'a')
			c--
		}
	}
	return string(bs)
}
