package leetcode0841

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	m := make(map[int]struct{}, n)

	var enter func(i int)
	enter = func(i int) {
		if _, ok := m[i]; ok {
			return
		}
		m[i] = struct{}{}
		for _, key := range rooms[i] {
			enter(key)
		}
	}
	enter(0)
	return len(m) == n
}
