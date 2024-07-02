package longest_valid_parentheses

func longestValidParenthesesStack(s string) int {
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
