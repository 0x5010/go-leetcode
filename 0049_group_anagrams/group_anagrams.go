package leetcode0049

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}
	r := make(map[string][]string)
	for _, str := range strs {
		count := make([]int, 26)
		for _, b := range []byte(str) {
			count[b-'a']++
		}

		bs := make([]byte, 26*2)
		for i, c := range count {
			bs[2*i] = '#'
			bs[2*i+1] = byte(c) + '0'
		}
		key := string(bs)
		r[key] = append(r[key], str)
	}
	res := make([][]string, len(r))
	i := 0
	for _, l := range r {
		res[i] = l
		i++
	}
	return res
}
