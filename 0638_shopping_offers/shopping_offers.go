package leetcode0638

import "math"

func shoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(price)
	for i, offer := range special {
		count := 0
		for j, p := range price {
			count += offer[j] * p
		}
		if count <= offer[n] {
			if i < len(special)-1 {
				special = append(special[:i], special[i+1:]...)
			} else {
				special = special[:len(special)-1]
			}
		}
	}

	var dfs func() int
	dfs = func() int {
		res := math.MaxInt32
		hasSpecial := false
		for _, offer := range special {
			flag := true
			for i := 0; i < n; i++ {
				if needs[i] < offer[i] {
					flag = false
				}
			}
			if flag {
				hasSpecial = true
				for i := 0; i < n; i++ {
					needs[i] -= offer[i]
				}
				v := offer[n] + dfs()
				if v < res {
					res = v
				}
				for i := 0; i < n; i++ {
					needs[i] += offer[i]
				}
			}
		}
		if !hasSpecial {
			res = 0
			for i, p := range price {
				res += needs[i] * p
			}
		}
		return res
	}
	return dfs()
}
