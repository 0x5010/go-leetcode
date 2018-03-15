package leetcode0167

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int, len(numbers))
	for i, n := range numbers {
		if m[target-n] != 0 {
			return []int{m[target-n], i + 1}
		}
		m[n] = i + 1
	}
	return nil
}
