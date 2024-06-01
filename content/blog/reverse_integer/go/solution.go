package main

import (
	"math"
)

func reverse(x int) int {
	left := x
	res := 0

	for left != 0 {
		right := left % 10
		left = left / 10

		if res > math.MaxInt32/10 || res == math.MaxInt32/10 && right > 7 {
			return 0
		}
		if res < math.MinInt32/10 || res == math.MinInt32/10 && right < -7 {
			return 0
		}

		res = res*10 + right
	}

	return res
}
