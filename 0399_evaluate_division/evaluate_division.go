package leetcode0399

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := make(map[string]float64)
	root := make(map[string]string)

	var find func(string) string
	find = func(s string) string {
		if root[s] == s {
			return s
		}
		return find(root[s])
	}
	var get func(string) float64
	get = func(s string) float64 {
		r := root[s]
		res := m[s]
		if r == s {
			return res
		}
		return res * get(r)
	}

	for i, e := range equations {
		a, b, v := e[0], e[1], values[i]

		if _, ok := root[a]; !ok {
			root[a] = a
		}
		if _, ok := root[b]; !ok {
			root[b] = b
		}
		if _, ok := m[a]; !ok {
			m[a] = 1.0
		}
		if _, ok := m[b]; !ok {
			m[b] = 1.0
		}
		r1, r2 := find(a), find(b)
		root[r2] = r1
		m[r2] = m[a] * v / m[b]
	}

	res := make([]float64, len(queries))
	for i, q := range queries {
		res[i] = -1.0
		a, b := q[0], q[1]
		if _, ok := root[a]; !ok {
			continue
		}
		if _, ok := root[b]; !ok {
			continue
		}
		r1, r2 := find(a), find(b)
		if r1 == r2 {
			res[i] = get(b) / get(a)
		}
	}
	return res
}
