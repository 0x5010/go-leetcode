package leetcode0676

type MagicDictionary struct {
	keys []string
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{
		keys: []string{},
	}
}

/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string) {
	for _, key := range dict {
		this.keys = append(this.keys, key)
	}
}

/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
	n := len(word)
	for _, key := range this.keys {
		if len(key) != n || key == word {
			continue
		}
		tmp := n
		for i := 0; i < n; i++ {
			if key[i] == word[i] {
				tmp--
			}
		}
		if tmp == 1 {
			return true
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dict);
 * param_2 := obj.Search(word);
 */
