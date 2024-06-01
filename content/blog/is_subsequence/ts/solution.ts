export const isSubsequence = (s: string, t: string): boolean => {
    let pos = 0

    if (s === "") {
        return true
    }

    for (let i = 0; i < t.length; i++) {
        const char = t[i]

        if (pos === s.length) {
            return true
        }

        if (char === s[pos]) {
            pos++
        }
    }

    return pos === s.length
}