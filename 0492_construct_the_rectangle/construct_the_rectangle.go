package leetcode0492

func constructRectangle(area int) []int {
	for w := mySqrt(area); w > 1; w-- {
		if area%w == 0 {
			return []int{area / w, w}
		}
	}
	return []int{area, 1}
}

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	low := 1
	high := x
	for low < high {
		mid := (low + high) / 2
		sq := mid * mid
		if sq > x {
			high = mid
		} else if sq < x {
			low = mid + 1
		} else {
			return mid
		}
	}

	return high - 1
}
