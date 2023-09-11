package main

func main() {

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	prefix := make([]rune, 0)
	firstStr := []rune(strs[0])
	runesStrs := make([][]rune, 0, len(strs))

	for i := 1; i < len(strs); i++ {
		if len([]rune(strs[i])) == 0 {
			return ""
		}

		runesStrs = append(runesStrs, []rune(strs[i]))
	}

loop:
	for idx, char := range firstStr {
		for _, runesStr := range runesStrs {
			if idx >= len(runesStr) {
				break loop
			}

			if runesStr[idx] != char {
				break loop
			}
		}

		prefix = append(prefix, char)
	}

	return string(prefix)
}
