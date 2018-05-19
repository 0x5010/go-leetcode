package leetcode0433

import "container/list"

func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	}
	m := make(map[string]struct{})
	for _, b := range bank {
		m[b] = struct{}{}
	}
	charSet := []byte{'A', 'C', 'G', 'T'}
	level := 0
	visited := make(map[string]struct{})
	queue := list.New()
	queue.PushBack(start)
	visited[start] = struct{}{}

	for queue.Len() != 0 {
		for n := queue.Len(); n > 0; n-- {
			curr := queue.Remove(queue.Front()).(string)
			if curr == end {
				return level
			}
			currBytes := []byte(curr)
			for i, b := range currBytes {
				old := b
				for _, c := range charSet {
					currBytes[i] = c
					next := string(currBytes)
					if _, ok := visited[next]; !ok {
						if _, ok := m[next]; ok {
							visited[next] = struct{}{}
							queue.PushBack(next)
						}
					}
				}
				currBytes[i] = old
			}
		}
		level++
	}
	return -1
}
