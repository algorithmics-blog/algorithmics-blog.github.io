export const hammingWeightPopCount = (n: number): number => {
    let mask = 1
    let count = 0

    for (let i = 0; i < 32; i++) {
        if (n & mask) {
            count++
        }
        mask <<= 1
    }

    return count
};