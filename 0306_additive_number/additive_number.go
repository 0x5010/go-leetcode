package leetcode0306

import "strconv"

func isAdditiveNumber(num string) bool {
	var dfs func(string, string, string) bool
	dfs = func(n1, n2, n string) bool {
		if len(n1) > 1 && n1[0] == '0' || len(n2) > 1 && n2[0] == '0' {
			return false
		}
		num1, _ := strconv.Atoi(n1)
		num2, _ := strconv.Atoi(n2)
		num3 := num1 + num2
		n3 := strconv.Itoa(num3)
		if len(n3) > len(n) {
			return false
		} else if len(n3) == len(n) {
			if n3 == n {
				return true
			}
			return false
		} else if n3 != n[:len(n3)] {
			return false
		}
		return dfs(n2, n3, n[len(n3):])
	}

	n := len(num)
	for i := 1; i <= n/2; i++ {
		for j := 1; j <= (n-i)/2; j++ {
			if dfs(num[:i], num[i:i+j], num[i+j:]) {
				return true
			}
		}
	}
	return false
}
