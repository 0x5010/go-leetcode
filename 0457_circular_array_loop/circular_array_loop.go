package leetcode457

func circularArrayLoop(nums []int) bool {
	n := len(nums)

	getIndex := func(i int) int {
		index := i + nums[i]
		if index >= 0 {
			return index % n
		}
		return n + (index % n)
	}

	for i, num := range nums {
		if num == 0 {
			continue
		}
		j, k := i, getIndex(i)
		for nums[k]*num > 0 && nums[getIndex(k)]*num > 0 {
			if j == k {
				if j == getIndex(j) {
					break
				}
				return true
			}
			j, k = getIndex(j), getIndex(getIndex(k))
		}
		j = i
		for nums[j]*num > 0 {
			next := getIndex(j)
			nums[j] = 0
			j = next
		}
	}
	return false
}
