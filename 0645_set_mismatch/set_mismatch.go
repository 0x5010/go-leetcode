package leetcode0645

func findErrorNums(nums []int) []int {
	dup, mis := 0, 0
	for i, num := range nums {
		if num != i+1 {
			nums[i] = 0
			for {
				pop := nums[num-1]
				if pop == num {
					dup = num
					break
				}
				nums[num-1] = num
				if pop == 0 {
					break
				}
				num = pop
			}
		}
	}

	for i, num := range nums {
		if num == 0 {
			mis = i + 1
			break
		}
	}
	return []int{dup, mis}
}
