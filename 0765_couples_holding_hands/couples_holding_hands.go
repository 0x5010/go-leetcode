package leetcode0765

func minSwapsCouples(row []int) int {
	n := len(row)
	partner := make([]int, n)
	indexs := make([]int, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			partner[i] = i + 1
		} else {
			partner[i] = i - 1
		}
		indexs[row[i]] = i
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := partner[indexs[partner[row[i]]]]; i != j; j = partner[indexs[partner[row[i]]]] {
			row[i], row[j] = row[j], row[i]
			indexs[row[i]], indexs[row[j]] = indexs[row[j]], indexs[row[i]]
			res++
		}
	}
	return res
}
