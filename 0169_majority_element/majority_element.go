package leetcode0169

func majorityElement(nums []int) int {
	var n, count int
	for _, num := range nums {
		if count == 0 {
			count++
			n = num
		} else {
			if num == n {
				count++
			} else {
				count--
			}
		}
	}
	return n
}
