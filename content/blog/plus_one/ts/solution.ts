export const plusOne = (digits: number[]): number[] => {
    const res: number[] = []
    let temp = 0

    for (let i = digits.length - 1; i >= 0; i--) {
        const digit = digits[i]
        temp += digit

        if (i == digits.length - 1) {
            temp += 1
        }

        if (temp < 10) {
            res.push(temp)
            temp = 0
            continue
        }

        res.push(temp % 10)
        temp = 1
    }

    if (temp != 0) {
        res.push(temp)
    }

    for (let i = 0, j = res.length - 1; i < j; i++, j--) {
        const tmp = res[i]
        res[i] = res[j]
        res[j] = tmp
    }

    return res
}
