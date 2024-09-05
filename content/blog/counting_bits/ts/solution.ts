export const countBits = (n: number): number[] => {
    const res: number[] = []

    for (let i = 0; i <= n; i++) {
        res[i] = (res[i >> 1] ?? 0) + (i & 1)
    }

    return res
}
