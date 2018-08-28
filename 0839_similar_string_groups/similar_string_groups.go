package leetcode0839

func numSimilarGroups(A []string) int {
	n := len(A)
	f := newForest(n)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if similar(A[i], A[j]) {
				f.union(i, j)
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		if f.parent[i] == i {
			res++
		}
	}
	return res
}

func similar(a, b string) bool {
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
			if diff > 2 {
				return false
			}
		}
	}
	return diff < 3
}

type forest struct {
	parent []int
	weight []int
}

func newForest(n int) *forest {
	parent := make([]int, n)
	weight := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		weight[i] = 1
	}
	return &forest{parent, weight}
}

func (f *forest) root(x int) int {
	r := x
	for f.parent[r] != r {
		r = f.parent[r]
	}
	for f.parent[x] != x {
		p := f.parent[x]
		f.parent[x] = r
		x = p
	}
	return r
}

func (f *forest) union(x, y int) {
	r1 := f.root(x)
	r2 := f.root(y)
	if f.weight[r1] < f.weight[r2] {
		f.parent[r1] = r2
		f.weight[r2] += f.weight[r1]
	} else {
		f.parent[r2] = r1
		f.weight[r1] += f.weight[r2]
	}
}
