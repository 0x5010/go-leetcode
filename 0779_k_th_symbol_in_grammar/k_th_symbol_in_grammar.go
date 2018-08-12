package leetcode0779

func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}
	if K&1 == 1 {
		return kthGrammar(N-1, (K+1)/2)
	}
	return (kthGrammar(N-1, K/2) + 1) & 1
}
