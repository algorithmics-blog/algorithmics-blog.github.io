export const reverseVowels = (s: string): string => {
    const dict = new Set(['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'])

    let leftIdx = 0
    let rightIdx = s.length - 1

    const chars = [...s]

    while (leftIdx < rightIdx) {
        const left = chars[leftIdx]
        const right = chars[rightIdx]

        if (!dict.has(left)) {
            leftIdx++
        } else if (!dict.has(right)) {
            rightIdx--
        } else {
            const tmp = chars[leftIdx]
            chars[leftIdx] = chars[rightIdx]
            chars[rightIdx] = tmp
            leftIdx++
            rightIdx--
        }
    }

    return chars.join('')
}
