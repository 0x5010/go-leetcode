package leetcode0476

func findComplement(num int) int {
	mask := ^0
	for num&mask != 0 {
		mask <<= 1
	}
	return ^mask & ^num
}
