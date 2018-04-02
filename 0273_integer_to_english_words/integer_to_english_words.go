package leetcode0273

import (
	"strings"
)

var (
	lessThan20 = []string{
		"",
		"One",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Eleven",
		"Twelve",
		"Thirteen",
		"Fourteen",
		"Fifteen",
		"Sixteen",
		"Seventeen",
		"Eighteen",
		"Nineteen",
		"Twenty",
	}

	tens = []string{
		"",
		"",
		"Twenty",
		"Thirty",
		"Forty",
		"Fifty",
		"Sixty",
		"Seventy",
		"Eighty",
		"Ninety",
	}

	thousands = []string{
		"",
		"Thousand",
		"Million",
		"Billion",
	}
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	res := ""
	for i := 0; num > 0; i++ {
		if num%1000 != 0 {
			res = lessK(num%1000) + thousands[i] + " " + res
		}
		num /= 1000
	}
	return strings.TrimRight(res, " ")
}

func lessK(num int) string {
	if num == 0 {
		return ""
	}

	if num <= 20 {
		return lessThan20[num] + " "
	}

	if num < 100 {
		return tens[num/10] + " " + lessK(num%10)
	}

	return lessThan20[num/100] + " Hundred " + lessK(num%100)
}
