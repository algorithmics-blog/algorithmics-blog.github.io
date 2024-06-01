export const removeStars = (s: string): string => {
    let res: string[] = []

    for (let char of s) {
        if (char === "*") {
            res.pop()
        } else {
            res.push(char)
        }
    }

    return res.join("")
}
