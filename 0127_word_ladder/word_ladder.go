package leetcode0127

var (
	exists = struct{}{}
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet, forward, backward := set(len(wordList)), set(1), set(1)
	for _, w := range wordList {
		wordSet[w] = exists
	}
	if _, ok := wordSet[endWord]; !ok {
		return 0
	}
	delete(wordSet, beginWord)
	forward[beginWord] = exists
	backward[endWord] = exists

	count := 2
	for len(backward) > 0 && len(forward) > 0 {
		if len(backward) < len(forward) {
			forward, backward = backward, forward
		}
		next := set(0)
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
						return count
					}
					if _, ok := wordSet[s]; ok {
						next[s] = exists
						delete(wordSet, s)
					}
				}
				word[i] = c
			}
		}
		forward = next
		count++
	}
	return 0
}

func set(n int) map[string]struct{} {
	return make(map[string]struct{}, n)
}
