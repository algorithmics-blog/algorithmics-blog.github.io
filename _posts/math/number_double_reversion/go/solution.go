package main

func reverse(num int) int {
	res := 0

	for num != 0 {
		right := num % 10
		num = num / 10

		res = res*10 + right
	}

	return res
}

func isSameAfterReversals(num int) bool {
	first := reverse(num)
	second := reverse(first)

	return num == second
}
