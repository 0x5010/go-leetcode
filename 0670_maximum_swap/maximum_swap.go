package leetcode0670

import "strconv"

func maximumSwap(num int) int {
	bs := []byte(strconv.Itoa(num))
	bucket := [10]int{}
	for i, b := range bs {
		bucket[b-'0'] = i
	}
	for i, b := range bs {
		for k := 9; k > int(b-'0'); k-- {
			j := bucket[k]
			if j > i {
				bs[i] = bs[j]
				bs[j] = b
				num, _ = strconv.Atoi(string(bs))
				return num
			}
		}
	}
	return num
}
