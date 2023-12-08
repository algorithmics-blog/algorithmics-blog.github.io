package main

import "math"

func myAtoi(s string) int {
	digitsMap := map[rune]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}

	res := 0
	trimmed := false
	isSignSet := false
	isPositive := true
	hasDigit := false

	for _, char := range s {
		if char == ' ' && !trimmed {
			continue
		}

		if !trimmed {
			trimmed = true
		}

		if hasDigit && (char == '-' || char == '+') {
			break
		}

		if !isSignSet && char == '-' {
			isPositive = false
			isSignSet = true
			continue
		}

		if !isSignSet && char == '+' {
			isPositive = true
			isSignSet = true
			continue
		}

		digit, ok := digitsMap[char]
		if !ok {
			break
		}

		hasDigit = true
		res = res*10 + digit

		if res > math.MaxInt32 {
			res = math.MaxInt32
			if !isPositive {
				res = res + 1
			}
			break
		}
	}

	if !isPositive {
		res = res * -1
	}

	return res
}
