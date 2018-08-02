package leetcode0721

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	f := newForest(len(accounts))
	commonEmails := make(map[string][]int)
	for i, account := range accounts {
		for _, email := range account[1:] {
			commonEmails[email] = append(commonEmails[email], i)
		}
	}
	for _, ids := range commonEmails {
		for i := 1; i < len(ids); i++ {
			f.union(ids[i-1], ids[i])
		}
	}
	m := make(map[int]map[string]struct{})
	for i, account := range accounts {
		id := f.root(i)
		emails, ok := m[id]
		if !ok {
			emails = make(map[string]struct{})
			m[id] = emails
		}
		for _, email := range account[1:] {
			emails[email] = struct{}{}
		}
	}
	res := make([][]string, 0, len(m))
	for id, emails := range m {
		tmp := make([]string, 0, len(emails)+1)
		tmp = append(tmp, accounts[id][0])
		for email := range emails {
			tmp = append(tmp, email)
		}
		sort.Strings(tmp[1:])
		res = append(res, tmp)
	}
	return res
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
