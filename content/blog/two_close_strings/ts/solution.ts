const fillCounter = (word: string): Record<string, number> => {
    const counter: Record<string, number> = {}

    for (let i = 0; i < word.length; i++) {
        if (!counter[word[i]]) {
            counter[word[i]] = 0
        }
        counter[word[i]]++
    }

    return counter
}

export const closeStrings = (word1: string, word2: string): boolean => {
    if (word1.length !== word2.length) {
        return false
    }

    const counter1 = fillCounter(word1)
    const counter2 = fillCounter(word2)
    const tempCounter: Record<number, number> = {}

    for (const [k, v] of Object.entries(counter1)) {
        const v2 = counter2[k]

        if (!v2) {
            return false
        }

        if (!tempCounter[v]) {
            tempCounter[v] = 0
        }
        if (!tempCounter[v2]) {
            tempCounter[v2] = 0
        }

        tempCounter[v]++
        tempCounter[v2]--
    }

    for (const [_, count] of Object.entries(tempCounter)) {
        if (count !== 0) {
            return false
        }
    }

    return true
}
