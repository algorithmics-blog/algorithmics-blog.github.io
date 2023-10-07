package main

type parenthesisWithCounter struct {
	str              string // генерированная строка
	generalOpenCount int    // суммарное кол-во открытых скобок в строке (не должно превышать n)
	notClosedCount   int    // кол-во незакрытых скобок в строке
}

func generateParenthesis(n int) []string {
	if n < 1 {
		return nil
	}

	parenthesisWithCounters := make([]parenthesisWithCounter, 0, 100)
	parenthesisWithCounters = append(parenthesisWithCounters, parenthesisWithCounter{
		str:              "(",
		generalOpenCount: 1,
		notClosedCount:   1,
	})

	for i := 1; i < n*2; i++ {
		newArray := make([]parenthesisWithCounter, 0, len(parenthesisWithCounters)*2)
		for _, item := range parenthesisWithCounters {
			if possibleToAddParenthesis('(', item, n) {
				newArray = append(newArray, parenthesisWithCounter{
					str:              item.str + "(",
					generalOpenCount: item.generalOpenCount + 1,
					notClosedCount:   item.notClosedCount + 1,
				})
			}

			if possibleToAddParenthesis(')', item, n) {
				newArray = append(newArray, parenthesisWithCounter{
					str:              item.str + ")",
					generalOpenCount: item.generalOpenCount,
					notClosedCount:   item.notClosedCount - 1,
				})
			}
		}

		parenthesisWithCounters = newArray
	}

	res := make([]string, 0, len(parenthesisWithCounters))
	for _, item := range parenthesisWithCounters {
		res = append(res, item.str)
	}

	return res
}

func possibleToAddParenthesis(parenthesis rune, prev parenthesisWithCounter, maxOpenCount int) bool {
	if parenthesis == '(' {
		if prev.generalOpenCount >= maxOpenCount {
			return false
		}

		return true
	}

	if parenthesis == ')' {
		if prev.notClosedCount <= 0 {
			return false
		}

		return true
	}

	return false
}

func generateParenthesisBruteforce(n int) []string {
	if n < 1 {
		return nil
	}

	combinations := generateAllPossibleParenthesis(n)
	res := make([]string, 0, len(combinations))
	for _, parenthesis := range combinations {
		if !isValid(parenthesis) {
			continue
		}

		res = append(res, parenthesis)
	}

	return res
}

func generateAllPossibleParenthesis(n int) []string {
	combinations := make([]string, 0, 100)
	combinations = append(combinations, "(")

	for i := 1; i < n*2; i++ {
		newArray := make([]string, 0, len(combinations))
		for _, item := range combinations {
			newArray = append(newArray, item+"(")
			newArray = append(newArray, item+")")
		}

		combinations = newArray
	}

	return combinations
}

func isValid(brackets string) bool {
	stack := make([]rune, 0, len(brackets))

	closeBracket := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range brackets {
		if openBracket, isClose := closeBracket[char]; isClose {
			if len(stack) == 0 {
				return false
			}

			lastOpenBracket := stack[len(stack)-1]

			if lastOpenBracket != openBracket {
				return false
			}

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
