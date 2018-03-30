package leetcode0242

// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	counts := [26]int{}
// 	for i := 0; i < len(s); i++ {
// 		counts[int(s[i]-'a')]++
// 		counts[int(t[i]-'a')]--
// 	}
// 	for _, c := range counts {
// 		if c != 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// follow up

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := [256]rune{}
	for _, r := range s {
		counts[r]++
	}
	for _, r := range t {
		counts[r]--
	}
	for _, c := range counts {
		if c != 0 {
			return false
		}
	}
	return true
}
