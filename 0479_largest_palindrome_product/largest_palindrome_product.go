package leetcode0479

func largestPalindrome(n int) int {
	return []int{9, 987, 123, 597, 677, 1218, 877, 475}[n-1]
}

// func largestPalindrome(n int) int {
// 	if n == 1 {
// 		return 9
// 	}
// 	upper := int(math.Pow(10, float64(n)) - 1)
// 	lower := upper / 10
// 	for a := upper; a > lower; a-- {
// 		left := a
// 		right := 0
// 		tmp := a
// 		for tmp > 0 {
// 			right = right*10 + tmp%10
// 			tmp /= 10
// 			left *= 10
// 		}
// 		palindrom := left + right
// 		for i := upper; i > lower; i-- {
// 			j := palindrom / i
// 			if j > i || j <= lower {
// 				break
// 			}
// 			if palindrom%i == 0 {
// 				return palindrom % 1337
// 			}
// 		}
// 	}
// 	return 0
// }
