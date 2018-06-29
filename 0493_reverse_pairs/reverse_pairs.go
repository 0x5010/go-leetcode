package leetcode0493

func reversePairs(nums []int) int {
	var recur func(int, int) int
	recur = func(l, r int) int {
		if l+1 >= r {

			return 0
		}

		m := (l + r) / 2
		res := recur(l, m) + recur(m, r)

		i, j := l, m
		for i < m && j < r {
			if nums[i] > 2*nums[j] {
				res += m - i
				j++
			} else {
				i++
			}
		}

		copy(nums[l:r], merge(nums[l:m], nums[m:r]))

		return res
	}
	return recur(0, len(nums))
}

func merge(a, b []int) []int {
	lenA, lenB := len(a), len(b)
	res := make([]int, lenA+lenB)

	var i, j, k int
	for i < lenA && j < lenB {
		if a[i] < b[j] {
			res[k] = a[i]
			i++
		} else {
			res[k] = b[j]
			j++
		}
		k++
	}

	if i == lenA {
		copy(res[k:], b[j:])
	} else {
		copy(res[k:], a[i:])
	}
	return res
}
