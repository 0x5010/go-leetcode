package leetcode0609

import "strings"

func findDuplicate(paths []string) [][]string {
	m := make(map[string][]string)
	for _, path := range paths {
		fs := strings.Split(path, " ")
		p := fs[0]
		for i := 1; i < len(fs); i++ {
			filename, content := splitFileNameAndContent(fs[i])
			m[content] = append(m[content], p+"/"+filename)
		}
	}

	res := [][]string{}
	for _, files := range m {
		if len(files) > 1 {
			res = append(res, files)
		}
	}
	return res
}

func splitFileNameAndContent(s string) (string, string) {
	i := strings.IndexByte(s, '(')
	return s[:i], s[i+1 : len(s)-1]
}
