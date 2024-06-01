export const decodeString = (s: string): string => {
    let res = ""

    for (let i = 0; i < s.length; i++) {
        if (s[i] >= 'a' && s[i] <= 'z') {
            res += s[i]
            continue
        }

        let count = 0
        while (s[i] != '[') {
            count = count*10 + Number(s[i])
            i++
        }

        let bracket = 1
        let j = i + 1
        while (bracket > 0) {
            if (s[j] == '[') {
                bracket++
            }

            if (s[j] == ']') {
                bracket--
            }

            j++
        }

        res += decodeString(s.substring(i+1, j-1)).repeat(count)

        i = j - 1
    }

    return res
}