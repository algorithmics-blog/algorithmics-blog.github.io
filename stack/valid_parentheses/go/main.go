package main

func main() {

}

func isValid(parenthesis string) bool {
	stack := make([]rune, 0, len(parenthesis))

	openParenthesis := map[rune]struct{}{
		'(': {},
		'[': {},
		'{': {},
	}

	closeParenthesis := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range parenthesis {
		if _, isOpen := openParenthesis[char]; isOpen {
			stack = append(stack, char)
			continue
		}

		if openChar, isClose := closeParenthesis[char]; isClose {
			if len(stack) == 0 {
				return false
			}
			lastOpenParenthesis := stack[len(stack)-1]

			if lastOpenParenthesis != openChar {
				return false
			}

			stack = stack[:len(stack)-1]
			continue
		}

		return false
	}

	return len(stack) == 0
}
