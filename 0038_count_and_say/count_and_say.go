package leetcode0038

import "strings"

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}
	res := "1"
	if n == 1 {
		return res
	}
	for i := 1; i < n; i++ {
		b := strings.Builder{}
		for j, count := 0, 1; j < len(res); j++ {
			if j+1 == len(res) || res[j] != res[j+1] {
				b.WriteByte(byte(count + '0'))
				b.WriteByte(res[j])
				count = 1
			} else {
				count++
			}
		}
		res = b.String()
	}
	return res
}
