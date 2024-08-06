export const hammingWeight = (n: number): number => {
    let count = 0

    while (n != 0) {
        count++
        n &= n - 1
    }

    return count
};