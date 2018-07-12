package leetcode0561

func arrayPairSum(nums []int) int {
	var buckets [20001]int8
	for _, num := range nums {
		buckets[num+10000]++
	}

	res := 0
	flag := true
	for num, count := range buckets {
		for count > 0 {
			if flag {
				res += num - 10000
			}
			flag = !flag
			count--
		}
	}

	return res
}
