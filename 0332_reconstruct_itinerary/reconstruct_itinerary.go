package leetcode0332

import "sort"

func findItinerary(tickets [][]string) []string {
	n := len(tickets)
	nexts := make(map[string][]string, n+1)
	for _, t := range tickets {
		nexts[t[0]] = append(nexts[t[0]], t[1])
	}

	for k := range nexts {
		sort.Strings(nexts[k])
	}

	res := make([]string, 0, n+1)

	var visit func(string)
	visit = func(airport string) {
		for len(nexts[airport]) != 0 {
			next := nexts[airport][0]
			nexts[airport] = nexts[airport][1:]
			visit(next)
		}
		res = append(res, airport)
	}

	visit("JFK")
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}
