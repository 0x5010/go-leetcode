package leetcode0043

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	len1, len2 := len(num1), len(num2)
	product := make([]int, len1+len2)

	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			product[i+j+1] += int(num1[i]-'0') * int(num2[j]-'0')
		}
	}

	for i := len(product) - 1; i > 0; i-- {
		product[i-1] += product[i] / 10
		product[i] = product[i] % 10
	}

	if product[0] == 0 {
		product = product[1:]
	}

	bs := make([]byte, len(product))
	for i, v := range product {
		bs[i] = '0' + byte(v)
	}
	return string(bs)
}
