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
