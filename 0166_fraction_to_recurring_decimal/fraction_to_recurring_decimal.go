package leetcode0166

import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	res := ""
	neg := false
	if numerator < 0 {
		numerator *= -1
		neg = !neg
	}
	if denominator < 0 {
		denominator *= -1
		neg = !neg
	}
	if neg {
		res = "-"
	}
	res += strconv.Itoa(numerator / denominator)
	rem := numerator % denominator
	if rem == 0 {
		return res
	}
	res += "."
	floatBytes := []byte{}
	itob := []byte("0123456789")
	indexMap := map[int]int{}
	for rem != 0 {
		if index, ok := indexMap[rem]; ok {
			return res + string(floatBytes[:index]) + "(" + string(floatBytes[index:]) + ")"
		}
		indexMap[rem] = len(floatBytes)
		rem *= 10
		b := itob[rem/denominator]
		rem %= denominator
		floatBytes = append(floatBytes, b)
	}
	return res + string(floatBytes)
}
