package leetcode0752

func openLock(deadends []string, target string) int {
	begin := map[[4]int]struct{}{}
	end := map[[4]int]struct{}{}
	deads := map[[4]int]struct{}{}
	for _, dead := range deadends {
		deads[convert(dead)] = struct{}{}
	}
	begin[convert("0000")] = struct{}{}
	end[convert(target)] = struct{}{}

	level := 0
	for len(begin) != 0 && len(end) != 0 {
		if len(begin) > len(end) {
			begin, end = end, begin
		}
		tmp := map[[4]int]struct{}{}
		for s := range begin {
			if _, ok := end[s]; ok {
				return level
			}
			if _, ok := deads[s]; ok {
				continue
			}
			for i := 0; i < 4; i++ {
				t1, t2 := s, s
				t1[i] = (t1[i] + 1) % 10
				t2[i] = (t2[i] + 9) % 10
				if _, ok := deads[t1]; !ok {
					tmp[t1] = struct{}{}
				}
				if _, ok := deads[t2]; !ok {
					tmp[t2] = struct{}{}
				}
			}
		}
		level++
		begin = tmp
	}
	return -1
}

func convert(s string) [4]int {
	res := [4]int{}
	for i := range res {
		res[i] = int(s[i] - '0')
	}
	return res
}
