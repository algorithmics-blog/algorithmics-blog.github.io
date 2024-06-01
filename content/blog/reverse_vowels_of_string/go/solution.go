package reverse_vowels_of_string

func reverseVowels(s string) string {
	dict := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}

	leftIdx := 0
	rightIdx := len(s) - 1

	sRunes := []rune(s)

	for leftIdx < rightIdx {
		left := sRunes[leftIdx]
		right := sRunes[rightIdx]

		if dict[left] == false {
			leftIdx++
		} else if dict[right] == false {
			rightIdx--
		} else {
			sRunes[leftIdx], sRunes[rightIdx] = sRunes[rightIdx], sRunes[leftIdx]
			leftIdx++
			rightIdx--
		}
	}

	return string(sRunes)
}
