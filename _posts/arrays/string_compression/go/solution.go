package main

import "strconv"

func compress(chars []byte) int {
	newPositionIdx, firstElemIdx := 0, 0

	for ; firstElemIdx < len(chars); firstElemIdx++ {
		lastChar := chars[firstElemIdx]

		lastElemIdx := firstElemIdx + 1
		for lastElemIdx < len(chars) && chars[lastElemIdx] == lastChar {
			lastElemIdx++
		}

		chars[newPositionIdx] = chars[lastElemIdx-1]
		newPositionIdx++

		if counter := lastElemIdx - firstElemIdx; counter > 1 {
			for _, countChar := range strconv.Itoa(counter) {
				chars[newPositionIdx] = byte(countChar)
				newPositionIdx++
			}

		}

		firstElemIdx = lastElemIdx - 1
	}

	chars = chars[:newPositionIdx]
	return len(chars)
}
