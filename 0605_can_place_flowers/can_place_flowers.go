package leetcode0605

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	m := len(flowerbed)
	for i, b := range flowerbed {
		if b == 0 &&
			((i+1 < m && flowerbed[i+1] == 0) || i+1 >= m) &&
			((i-1 >= 0 && flowerbed[i-1] == 0) || i-1 < 0) {
			flowerbed[i] = 1
			n--
			if n == 0 {
				return true
			}
		}
	}
	return false
}
