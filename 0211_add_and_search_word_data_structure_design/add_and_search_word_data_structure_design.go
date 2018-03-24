package leetcode0211

type Trie struct {
	children [26]*Trie
	isLeaf   bool
}

type WordDictionary struct {
	root *Trie
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		root: &Trie{},
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	p := this.root
	for i := 0; i < len(word); i++ {
		ci := int(word[i] - 'a')
		if p.children[ci] == nil {
			p.children[ci] = &Trie{}
		}
		p = p.children[ci]
	}
	p.isLeaf = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return search(word, 0, this.root)
}

func search(word string, i int, node *Trie) bool {
	if i == len(word) {
		return node.isLeaf
	}

	if word[i] == '.' {
		for j := 0; j < 26; j++ {
			if node.children[j] != nil && search(word, i+1, node.children[j]) {
				return true
			}
		}
	} else {
		ci := int(word[i] - 'a')
		return node.children[ci] != nil && search(word, i+1, node.children[ci])
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
