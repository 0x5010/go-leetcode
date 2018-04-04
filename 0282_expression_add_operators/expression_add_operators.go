package leetcode0282

import "strconv"

func addOperators(num string, target int) []string {
	n := len(num)
	res := []string{}
	var dfs func(string, int, int, int)
	dfs = func(path string, index, result, pre int) {
		if index == n && result == target {
			res = append(res, path)
			return
		}

		for i := index; i < n; i++ {
			if num[index] == '0' && i != index {
				return
			}
			currStr := num[index : i+1]
			curr, _ := strconv.Atoi(currStr)
			if index == 0 {
				dfs(currStr, i+1, curr, curr)
			} else {
				dfs(path+"+"+currStr, i+1, result+curr, curr)
				dfs(path+"-"+currStr, i+1, result-curr, -curr)
				dfs(path+"*"+currStr, i+1, result-pre+pre*curr, pre*curr)
			}
		}
	}
	dfs("", 0, 0, 0)
	return res
}
