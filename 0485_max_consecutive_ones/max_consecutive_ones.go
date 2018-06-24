package leetcode0485

func findMaxConsecutiveOnes(nums []int) int {
	res := 0
	count := 0
	for _, num := range nums {
		if num == 1 {
			count++
		} else {
			if count > res {
				res = count
			}
			count = 0
		}
	}
	if count > res {
		return count
	}
	return res
}
