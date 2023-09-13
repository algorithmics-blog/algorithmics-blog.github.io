package main

func main() {

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := make([]rune, 0)
	firstStr := strs[0]

	for idx, char := range firstStr {
		for _, str := range strs {
			if idx >= len(str) || str[idx] != firstStr[idx] {
				return string(prefix)
			}
		}

		prefix = append(prefix, char)
	}

	return string(prefix)
}
