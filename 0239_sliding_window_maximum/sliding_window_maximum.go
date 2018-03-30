package leetcode0239

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	indexs := []int{}
	for i, num := range nums {
		for len(indexs) != 0 && indexs[0] <= i-k {
			indexs = indexs[1:]
		}
		j := len(indexs) - 1
		for j >= 0 && nums[indexs[j]] <= num {
			j--
		}
		indexs = append(indexs[:j+1], i)
		if i >= k-1 {
			res = append(res, nums[indexs[0]])
		}
	}
	return res
}
