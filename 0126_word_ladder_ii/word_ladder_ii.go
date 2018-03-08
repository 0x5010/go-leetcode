package leetcode0126

var (
	exists = struct{}{}
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	paths := [][]string{}
	path := []string{beginWord}
	if beginWord == endWord {
		paths = append(paths, path)
		return paths
	}

	wordSet, forward, backward := set(len(wordList)), set(1), set(1)

	for _, w := range wordList {
		wordSet[w] = exists
	}
	if _, ok := wordSet[endWord]; !ok {
		return paths
	}
	forward[beginWord] = exists
	backward[endWord] = exists
	tree := make(map[string][]string)

	var dfs func(string, string, map[string][]string, []string)
	dfs = func(
		beginWord, endWord string,
		tree map[string][]string,
		path []string,
	) {
		if beginWord == endWord {
			tmp := make([]string, len(path))
			copy(tmp, path)
			paths = append(paths, tmp)
			return
		}
		for _, it := range tree[beginWord] {
			path = append(path, it)
			dfs(it, endWord, tree, path)
			path = path[:len(path)-1]
		}
	}

	if buildTree(wordSet, forward, backward, tree, false) {
		dfs(beginWord, endWord, tree, path)
	}
	return paths
}

func set(n int) map[string]struct{} {
	return make(map[string]struct{}, n)
}

func buildTree(
	wordSet, forward, backward map[string]struct{},
	tree map[string][]string,
	reversed bool,
) bool {
	if len(forward) == 0 {
		return false
	}
	if len(forward) > len(backward) {
		return buildTree(wordSet, backward, forward, tree, !reversed)
	}
	for k := range forward {
		delete(wordSet, k)
	}
	for k := range backward {
		delete(wordSet, k)
	}
	next := set(0)
	done := false
	for it := range forward {
		word := []byte(it)
		for i := 0; i < len(word); i++ {
			c := word[i]
			for word[i] = 'a'; word[i] <= 'z'; word[i]++ {
				if word[i] == c {
					continue
				}
				s := string(word)
				if _, ok := backward[s]; ok {
					done = true
					if reversed {
						tree[s] = append(tree[s], it)
					} else {
						tree[it] = append(tree[it], s)
					}
				} else if !done {
					if _, ok := wordSet[s]; ok {
						next[s] = exists
						if reversed {
							tree[s] = append(tree[s], it)
						} else {
							tree[it] = append(tree[it], s)
						}
					}
				}
			}
			word[i] = c
		}
	}
	return done || buildTree(wordSet, next, backward, tree, reversed)
}
