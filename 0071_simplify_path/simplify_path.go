package leetcode0071

import (
	"strings"
)

func simplifyPath(path string) string {
	comps := strings.Split(path, "/")

	res := []string{}
	for _, s := range comps {
		switch s {
		case ".", "":
			// do nothing
		case "..":
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		default:
			res = append(res, s)
		}
	}
	return "/" + strings.Join(res, "/")
}
