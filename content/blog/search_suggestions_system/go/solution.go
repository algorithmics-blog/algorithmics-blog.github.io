package search_suggestions_system

import (
	"sort"
	"strings"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	res := make([][]string, 0, len(searchWord))

	var idx int
	items := products
	for idx = range searchWord {
		items, _ = suggestions(items, searchWord[:idx+1])
		if len(items) == 0 {
			break
		}

		maxLen := 3
		if len(items) < maxLen {
			maxLen = len(items)
		}
		res = append(res, items[:maxLen])
	}

	if idx < len(searchWord)-1 {
		res = gd(res, len(searchWord)-idx)
	}

	return res
}

func suggestions(products []string, prefix string) ([]string, bool) {
	startIdx := findFirstProductWithPrefix(products, prefix)
	if startIdx == -1 {
		return []string{}, true
	}

	res := make([]string, 0, 3)

	for i := startIdx; i < len(products); i++ {
		if !strings.HasPrefix(products[i], prefix) {
			break
		}

		res = append(res, products[i])
	}

	return res, false
}

func findFirstProductWithPrefix(products []string, prefix string) int {
	strWithPrefixIdx := -1
	startIdx, endIdx := 0, len(products)-1

	for startIdx <= endIdx {
		middleIdx := (endIdx + startIdx) / 2

		if strings.HasPrefix(products[middleIdx], prefix) {
			strWithPrefixIdx = middleIdx
			break
		}
		productRunes := []rune(products[middleIdx])

		for idx, char := range prefix {
			if idx > len(productRunes)-1 {
				startIdx = middleIdx + 1
				break
			}

			if productRunes[idx] == char {
				continue
			}

			if productRunes[idx] < char {
				startIdx = middleIdx + 1
			} else {
				endIdx = middleIdx - 1
			}

			break
		}

	}

	res := -1
	for i := strWithPrefixIdx; i >= 0; i-- {
		if strings.HasPrefix(products[i], prefix) {
			res = i
			continue
		}

		break
	}

	return res
}

func gd(res [][]string, tailLength int) [][]string {
	for i := 0; i < tailLength; i++ {
		res = append(res, []string{})
	}

	return res
}
