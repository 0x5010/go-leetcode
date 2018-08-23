package leetcode0014

import "strings"

type trieNode struct {
	links []*trieNode
	isEnd bool
	size  int
}

func newTrieNode() *trieNode {
	return &trieNode{
		links: make([]*trieNode, 26),
	}
}

func (tn *trieNode) get(ch byte) *trieNode {
	return tn.links[ch-'a']
}

func (tn *trieNode) put(ch byte, node *trieNode) {
	tn.links[ch-'a'] = node
	tn.size++
}

func (tn *trieNode) containsKey(ch byte) bool {
	return tn.get(ch) != nil
}

type trie struct {
	root *trieNode
}

func newTrie() *trie {
	return &trie{root: newTrieNode()}
}

func (t *trie) insert(s string) {
	node := t.root
	for _, ch := range []byte(s) {
		if !node.containsKey(ch) {
			node.put(ch, newTrieNode())
		}
		node = node.get(ch)
	}
	node.isEnd = true
}

func (t *trie) searchLongestPrefix(s string) string {
	node := t.root
	prifix := strings.Builder{}
	for _, ch := range []byte(s) {
		if node.containsKey(ch) && node.size == 1 && !node.isEnd {
			prifix.WriteByte(ch)
			node = node.get(ch)
		} else {
			break
		}
	}
	return prifix.String()
}

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	t := newTrie()
	for _, s := range strs {
		t.insert(s)
	}
	return t.searchLongestPrefix(strs[0])
}
