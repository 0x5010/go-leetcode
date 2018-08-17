package leetcode0808

func soupServings(N int) float64 {
	if N >= 4800 {
		return 1
	}
	n := (N + 24) / 25
	m := make([][]float64, n+1)
	for i := 0; i < n+1; i++ {
		m[i] = make([]float64, n+1)
	}
	var f func(int, int) float64
	f = func(a, b int) float64 {
		if a <= 0 && b <= 0 {
			return 0.5
		}
		if a <= 0 {
			return 1
		}
		if b <= 0 {
			return 0
		}
		if m[a][b] > 0 {
			return m[a][b]
		}
		m[a][b] = 0.25 * (f(a-4, b) + f(a-3, b-1) + f(a-2, b-2) + f(a-1, b-3))
		return m[a][b]
	}
	return f(n, n)
}
