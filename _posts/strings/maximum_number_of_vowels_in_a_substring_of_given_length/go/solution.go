package main

func maxVowels(s string, k int) int {
	vowelsCount := 0

	for rightIdx := 0; rightIdx < len(s) && rightIdx < k; rightIdx++ {
		if isVowel(s[rightIdx]) {
			vowelsCount++
		}
	}

	if vowelsCount == k {
		return vowelsCount
	}

	currentVowelsCount := vowelsCount
	leftIdx := 0

	for rightIdx := k; rightIdx < len(s); rightIdx++ {
		if isVowel(s[rightIdx]) {
			currentVowelsCount++
		}
		if isVowel(s[leftIdx]) {
			currentVowelsCount--
		}

		leftIdx++
		if currentVowelsCount > vowelsCount {
			vowelsCount = currentVowelsCount
			if vowelsCount == k {
				return vowelsCount
			}
		}
	}

	return vowelsCount
}

func isVowel(ch byte) bool {
	switch ch {
	case 'a':
		return true
	case 'e':
		return true
	case 'i':
		return true
	case 'o':
		return true
	case 'u':
		return true
	}
	return false
}
