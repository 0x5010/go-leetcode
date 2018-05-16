package leetcode0415

func addStrings(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	if m == 0 {
		return num2
	}
	if n == 0 {
		return num1
	}
	if m > n {
		m, n, num1, num2 = n, m, num2, num1
	}
	carry := byte(0)
	buf := make([]byte, n+1)
	for i := 1; i <= n; i++ {
		if i <= m {
			buf[n-i+1] = num1[m-i] - '0'
		}
		buf[n-i+1] += num2[n-i] + carry
		if buf[n-i+1] > '9' {
			buf[n-i+1] -= 10
			carry = byte(1)
		} else {
			carry = byte(0)
		}
	}
	if carry == 1 {
		buf[0] = '1'
		return string(buf)
	}
	return string(buf[1:])
}
