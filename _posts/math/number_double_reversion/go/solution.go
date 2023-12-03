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
	reversed1 := reverse(num)
	reversed2 := reverse(reversed1)

	return num == reversed2
}
