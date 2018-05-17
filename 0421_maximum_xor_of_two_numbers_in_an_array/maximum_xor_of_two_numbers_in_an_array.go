package leetcode0421

func findMaximumXOR(nums []int) int {
	res, mask := 0, 0

	for i := 31; i >= 0; i-- {
		mask |= 1 << uint(i)
		m := make(map[int]struct{})
		for _, num := range nums {
			m[num&mask] = struct{}{}
		}
		tmp := res | 1<<uint(i)
		for k := range m {
			if _, ok := m[tmp^k]; ok {
				res = tmp
				break
			}
		}
	}
	return res
}
