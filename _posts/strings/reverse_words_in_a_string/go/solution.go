package main

func reverseWords(s string) string {
	startIdx := 0
	bStr := []byte(s)

	for i := 0; i < len(bStr); i++ {
		if bStr[i] == byte(' ') {
			reverseWord(bStr, startIdx, i-1)

			startIdx = i + 1
			continue
		}
	}

	// корнер кейс - переворачиваем последнее слово в строке, так как, если строка не заканчивается на ' ',
	// мы не обработаем этот кейс в цикле выше
	reverseWord(bStr, startIdx, len(bStr)-1)

	reverseWord(bStr, 0, len(bStr)-1)

	return string(trimSpaces(bStr))
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
