package is_subsequence

func isSubsequence(s string, t string) bool {
	pos := 0
	runes := []rune(s)

	if s == "" {
		return true
	}

	for _, char := range []rune(t) {
		if pos == len(s) {
			return true
		}

		if char == runes[pos] {
			pos++
		}
	}

	return pos == len(s)
}
