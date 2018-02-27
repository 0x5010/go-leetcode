package leetcode0093

func restoreIpAddresses(s string) []string {
	n := len(s)
	res := []string{}
	var dfs func(int, int, string)
	dfs = func(start, step int, ip string) {
		if start == n && step == 4 {
			res = append(res, ip[:len(ip)-1])
			return
		}
		if n-start > (4-step)*3 || n-start < 4-step {
			return
		}
		num := 0
		for i := start; i < start+3 && i < n; i++ {
			num = num*10 + int(s[i]-'0')
			if num < 256 {
				ip += s[i : i+1]
				dfs(i+1, step+1, ip+".")
			}
			if num == 0 {
				break
			}
		}
	}
	dfs(0, 0, "")
	return res
}
