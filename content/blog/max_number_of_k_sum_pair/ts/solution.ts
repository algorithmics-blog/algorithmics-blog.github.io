export const maxOperations = (nums: number[], k: number): number => {
    const digits: Record<number, number> = {}
    let res = 0

    for (const num of nums) {
        const diff = k - num

        if (digits[diff] > 0) {
            digits[diff] = (digits[diff] || 0) - 1
            res++
        } else {
            digits[num] = (digits[num] || 0) + 1
        }
    }

    return res
}
