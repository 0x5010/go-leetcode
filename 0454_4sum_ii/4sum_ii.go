package leetcode454

func fourSumCount(A []int, B []int, C []int, D []int) int {
	res := 0
	sum := map[int]int{}
	for _, a := range A {
		for _, b := range B {
			sum[a+b]++
		}
	}
	for _, c := range C {
		for _, d := range D {
			res += sum[-(c + d)]
		}
	}
	return res
}
