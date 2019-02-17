package leetcode0952

func largestComponentSize(A []int) int {
	notPrime := [100001]bool{}
	primes := map[int]struct{}{}
	for i := 2; i < 100001; i++ {
		if !notPrime[i] {
			primes[i] = struct{}{}
			for j := 2; j*i < 100001; j++ {
				notPrime[j*i] = true
			}
		}
	}
	primesIdx := [100001]int{}
	for i := 0; i < 100001; i++ {
		primesIdx[i] = -1
	}
	n := len(A)
	f := newForest(n)
	for i, a := range A {
		for p := range primes {
			if _, ok := primes[a]; ok {
				p = a
			}
			if a%p == 0 {
				idx := primesIdx[p]
				if idx != -1 {
					f.union(idx, i)
				}
				primesIdx[p] = i
				for a%p == 0 {
					a /= p
				}
			}
			if a == 1 {
				break
			}
		}
	}
	return f.max
}

type forest struct {
	f       map[int]int
	islands int
}

func (f *forest) root(x int) int {
	if tmp, ok := f.f[x]; ok {
		if x != tmp {
			f.f[x] = f.root(tmp)
		}
	} else {
		f.f[x] = x
		f.islands++
	}
	return f.f[x]
}

func (f *forest) union(x, y int) {
	r1 := f.root(x)
	r2 := f.root(y)
	if f.f[r1] != r2 {
		f.f[r1] = r2
		f.islands--
	}
}
