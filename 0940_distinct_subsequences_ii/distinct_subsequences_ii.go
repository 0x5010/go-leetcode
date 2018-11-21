package leetcode0940

func distinctSubseqII(S string) int {
	mod := 1000000007
	end := [26]int{}
	res := 0
	for i := 0; i < len(S); i++ {
		c := S[i] - 'a'
		tmp := (res + 1 - end[c]) % mod
		res = (res + tmp) % mod
		end[c] = (end[c] + tmp) % mod
	}
	return (res + mod) % mod
}
