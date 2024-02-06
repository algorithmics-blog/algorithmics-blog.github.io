package main

func reverseWords(s string) string {
	startIdx := 0
	endIdx := 0

	trimmed := trimSpaces([]byte(s))

	for startIdx < len(trimmed) {
		for endIdx < len(trimmed) && trimmed[endIdx] != byte(' ') {
			endIdx++
		}

		reverseWord(trimmed, startIdx, endIdx-1)
		startIdx = endIdx + 1
		endIdx++
	}

	reverseWord(trimmed, 0, len(trimmed)-1)

	return string(trimmed)
}

func reverseWord(s []byte, i, j int) {
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func trimSpaces(bStr []byte) []byte {
	isStart := true
	j := 0
	for i := 0; i < len(bStr); i++ {
		if bStr[i] == byte(' ') && isStart {
			continue
		}

		bStr[j] = bStr[i]
		j++
		isStart = bStr[i] == byte(' ')
	}

	if bStr[j-1] == byte(' ') {
		j--
	}

	return bStr[:j]
}
