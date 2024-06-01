const reverseWord = (s: string, i: number, j: number): string => {
    const strArr = [...s]

    while (i < j) {
        const temp = strArr[i]
        strArr[i] = strArr[j]
        strArr[j] = temp
        i++
        j--
    }

    return strArr.join('')
}

const trimSpaces = (s: string): string => {
    let isStart = true
    let j = 0
    let strArr = [...s]

    for (let i = 0; i < strArr.length; i++) {
        if (strArr[i] == ' ' && isStart) {
            continue
        }

        strArr[j] = strArr[i]
        isStart = strArr[i] == ' '
        j++
    }

    if (strArr[j - 1] == ' ') {
        j--
    }

    return strArr
        .join('')
        .substring(0, j)
}

export const reverseWords = (s: string): string => {
    let startIdx = 0
    let endIdx = 0

    let trimmed = trimSpaces(s)

    while (startIdx < trimmed.length) {
        while (endIdx < trimmed.length && trimmed[endIdx] !== ' ') {
            endIdx++
        }

        trimmed = reverseWord(trimmed, startIdx, endIdx - 1)
        startIdx = endIdx + 1
        endIdx++
    }

    return reverseWord(trimmed, 0, trimmed.length - 1)
}
