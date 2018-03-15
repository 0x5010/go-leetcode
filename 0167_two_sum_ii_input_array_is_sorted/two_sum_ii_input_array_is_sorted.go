package leetcode0167

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	if n < 2 {
		return nil
	}
	i, j := 0, n-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
	return nil
}
