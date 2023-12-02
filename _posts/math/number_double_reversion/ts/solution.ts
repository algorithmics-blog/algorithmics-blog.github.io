const reverse = (num: number): number => {
    let res = 0

    while (num != 0) {
        const right = num % 10
        num = Math.floor(num / 10)
        res = res * 10 + right
    }

    return res
}

export const isSameAfterReversals = (num: number): boolean => {
    const first = reverse(num)
    const second = reverse(first)

    return num == second
}
