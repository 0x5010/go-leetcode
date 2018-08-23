package leetcode0038

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}
	res := "1"
	if n == 1 {
		return res
	}
	bs := []byte{}
	for i := 1; i < n; i++ {
		for j, count := 0, 1; j < len(res); j++ {
			if j+1 == len(res) || res[j] != res[j+1] {
				bs = append(bs, byte(count+'0'))
				bs = append(bs, res[j])
				count = 1
			} else {
				count++
			}
		}
		res = string(bs)
		bs = bs[:0]
	}
	return res
}
