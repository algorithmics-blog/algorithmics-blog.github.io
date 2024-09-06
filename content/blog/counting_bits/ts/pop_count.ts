const hammingWeight = (n: number): number => {
    let count = 0

    while (n != 0) {
        count++
        n &= n - 1
    }

    return count
}

export const countBitsPopCount = (n: number): number[] => {
    const res: number[] = []

    for (let i = 0; i <= n; i++) {
        res.push(hammingWeight(i))
    }

    return res
}
