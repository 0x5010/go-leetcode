package leetcode0495

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 || duration == 0 {
		return 0
	}

	res, start, end := 0, timeSeries[0], timeSeries[0]+duration
	for _, ts := range timeSeries {
		if ts > end {
			res += end - start
			start = ts
		}
		end = ts + duration
	}
	res += end - start
	return res
}
