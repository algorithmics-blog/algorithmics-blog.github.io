package longest_valid_parentheses

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
