package leetcode0208

type Trie struct {
	children [26]*Trie
	isLeaf   bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	p := this
	for i := 0; i < len(word); i++ {
		ci := int(word[i] - 'a')
		if p.children[ci] == nil {
			p.children[ci] = &Trie{}
		}
		p = p.children[ci]
	}
	p.isLeaf = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	p := this
	for i := 0; i < len(word); i++ {
		ci := int(word[i] - 'a')
		if p.children[ci] == nil {
			return false
		}
		p = p.children[ci]
	}
	return p.isLeaf
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	p := this
	for i := 0; i < len(prefix); i++ {
		ci := int(prefix[i] - 'a')
		if p.children[ci] == nil {
			return false
		}
		p = p.children[ci]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
