const isVowel = (char: string): boolean => {
    switch (char) {
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

export const maxVowels = (s: string, k: number): number => {
    let vowelsCount = 0

    for (let rightIdx = 0; rightIdx < s.length && rightIdx < k; rightIdx++) {
        if (isVowel(s[rightIdx])) {
            vowelsCount++
        }
    }

    if (vowelsCount == k) {
        return vowelsCount
    }

    let currentVowelsCount = vowelsCount
    let leftIdx = 0

    for (let rightIdx = k; rightIdx < s.length; rightIdx++) {
        if (isVowel(s[rightIdx])) {
            currentVowelsCount++
        }
        if (isVowel(s[leftIdx])) {
            currentVowelsCount--
        }

        leftIdx++
        if (currentVowelsCount > vowelsCount) {
            vowelsCount = currentVowelsCount
            if (vowelsCount == k) {
                return vowelsCount
            }
        }
    }

    return vowelsCount
}
