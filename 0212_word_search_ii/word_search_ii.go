package leetcode0212

func findWords(board [][]byte, words []string) []string {
	m := len(board)
	if m == 0 {
		return nil
	}
	n := len(board[0])

	root := buildTrie(words)
	res := []string{}
	var mark byte = '#'
	var dfs func(int, int, *Trie)
	dfs = func(i, j int, node *Trie) {
		if i < 0 || i > m-1 || j < 0 || j > n-1 {
			return
		}

		b := board[i][j]
		ci := int(b - 'a')
		if b == mark || node.children[ci] == nil {
			return
		}

		node = node.children[ci]
		if node.word != "" {
			res = append(res, node.word)
			node.word = ""
		}

		board[i][j] = mark
		dfs(i-1, j, node)
		dfs(i+1, j, node)
		dfs(i, j-1, node)
		dfs(i, j+1, node)
		board[i][j] = b
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, root)
		}
	}
	return res
}

type Trie struct {
	children [26]*Trie
	word     string
}

func buildTrie(words []string) *Trie {
	root := &Trie{}
	for _, word := range words {
		p := root
		for i := 0; i < len(word); i++ {
			ci := int(word[i] - 'a')
			if p.children[ci] == nil {
				p.children[ci] = &Trie{}
			}
			p = p.children[ci]
		}
		p.word = word
	}
	return root
}
