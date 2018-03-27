package leetcode0229

func majorityElement(nums []int) []int {
	candidate1, candidate2, count1, count2 := 0, 1, 0, 0
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		} else if count1 == 0 {
			candidate1, count1 = num, 1
		} else if count2 == 0 {
			candidate2, count2 = num, 1
		} else {
			count1--
			count2--
		}
	}
	res := []int{}
	if majority(nums, candidate1) {
		res = append(res, candidate1)
	}
	if majority(nums, candidate2) {
		res = append(res, candidate2)
	}
	return res
}

func majority(nums []int, num int) bool {
	count := 0
	for _, n := range nums {
		if n == num {
			count++
		}
	}
	return count > len(nums)/3
}
