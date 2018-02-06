package leetcode0017

var m = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
	'0': []string{" "},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	res := []string{""}
	for _, d := range []byte(digits) {
		tmp := []string{}
		for _, s := range res {
			for _, ch := range m[d] {
				tmp = append(tmp, s+ch)
			}
		}
		res = tmp
	}
	return res
}
