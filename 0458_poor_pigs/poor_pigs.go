package leetcode458

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	res := 0
	for pow((minutesToTest/minutesToDie+1), res) < buckets {
		res++
	}

	return res
}

func pow(x, n int) int {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	res := pow(x, n>>1)
	if n&1 == 0 {
		return res * res
	}
	return res * res * x
}
