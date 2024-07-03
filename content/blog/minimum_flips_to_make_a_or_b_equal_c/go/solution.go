package main

func minFlips(a int, b int, c int) int {
	res := 0

	for a > 0 || b > 0 || c > 0 {
		bitA := a % 2
		bitB := b % 2
		bitC := c % 2

		a = a / 2
		b = b / 2
		c = c / 2

		if bitA|bitB == bitC {
			// Условие задачи выполняется, значит биты на этой позиции менять не нужно
			continue
		}

		if bitC == 0 {
			res += bitA + bitB
		} else {
			res += 1
		}
	}

	return res
}
