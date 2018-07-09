package leetcode0537

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(a string, b string) string {
	r1, i1 := complex(a)
	r2, i2 := complex(b)
	return fmt.Sprintf("%d+%di", r1*r2-i1*i2, r1*i2+r2*i1)
}

func complex(s string) (int, int) {
	l := strings.Split(s, "+")
	r, _ := strconv.Atoi(l[0])
	i, _ := strconv.Atoi(l[1][:len(l[1])-1])
	return r, i
}
