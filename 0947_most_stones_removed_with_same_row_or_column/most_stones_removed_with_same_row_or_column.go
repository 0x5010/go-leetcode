package leetcode0947

func removeStones(stones [][]int) int {
	f := &forest{
		f: make(map[int]int),
	}
	for _, stone := range stones {
		f.union(stone[0], ^stone[1])
	}
	return len(stones) - f.islands
}

type forest struct {
	f       map[int]int
	islands int
}

func (f *forest) root(x int) int {
	if tmp, ok := f.f[x]; ok {
		if x != tmp {
			f.f[x] = f.root(tmp)
		}
	} else {
		f.f[x] = x
		f.islands++
	}
	return f.f[x]
}

func (f *forest) union(x, y int) {
	r1 := f.root(x)
	r2 := f.root(y)
	if f.f[r1] != r2 {
		f.f[r1] = r2
		f.islands--
	}
}
