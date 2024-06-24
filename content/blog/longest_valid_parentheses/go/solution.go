package main

func longestValidParentheses(s string) int {
	stack := make([]int, 0, len(s))
	stack = append(stack, -1)
	res := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}

		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}

		res = max(res, i-stack[len(stack)-1])
	}

	return res
}

// Неоптимальное решение через рекурсию и скользящее окно

func longestValidParenthesesWindowing(s string) int {
	return recursiveLongestValidParenthesesWindowing(s, len(s))
}

func recursiveLongestValidParenthesesWindowing(s string, substrLen int) int {
	if substrLen == 0 {
		return 0
	}
	left := 0
	right := substrLen
	for right <= len(s) {
		if isValidParenthesesByCounter(s[left:right]) {
			return right - left
		}
		left++
		right++
	}

	return recursiveLongestValidParenthesesWindowing(s, substrLen-1)
}

func isValidParenthesesByCounter(s string) bool {
	counter := 0
	for _, char := range s {
		if char == '(' {
			counter++
			continue
		}

		if counter == 0 {
			return false
		}

		counter--
	}

	return counter == 0
}
