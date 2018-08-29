package leetcode0842

/**
 * // This is the Master's API interface.
 * // You should not implement it, or speculate about its implementation
 * type Master struct {
 * }
 *
 * func (this *Master) Guess(word string) int {}
 */

func findSecretWord(wordlist []string, master *Master) {
	for i := 0; i < 10; i++ {

		if len(wordlist) < 1 {
			return
		}
		if len(wordlist) == 1 {
			master.Guess(wordlist[0])
			return
		}
		newlist := []string{}
		num := master.Guess(wordlist[1])
		if num == 6 {
			return
		}
		for i := 0; i < len(wordlist); i++ {
			if match(wordlist[i], wordlist[1]) == num {
				newlist = append(newlist, wordlist[i])
			}
		}
		wordlist = newlist
	}
}

func match(a, b string) int {
	count := 0
	for i := 0; i < 6; i++ {
		if a[i] == b[i] {
			count++
		}
	}
	return count
}
