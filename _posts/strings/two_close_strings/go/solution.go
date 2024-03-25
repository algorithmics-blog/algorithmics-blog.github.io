package two_close_strings

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	counter1 := make(map[rune]int)
	counter2 := make(map[rune]int)

	for _, char := range word1 {
		counter1[char]++
	}

	for _, char := range word2 {
		counter2[char]++
	}

	tempCounter := make(map[int]int)

	for k, v := range counter1 {
		v2, exist := counter2[k]
		if !exist {
			return false
		}

		tempCounter[v]++
		tempCounter[v2]--
	}

	for _, count := range tempCounter {
		if count != 0 {
			return false
		}
	}

	return true
}
