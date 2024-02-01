export const singleNumberMap = (nums: number[]): number => {
    const countMap: Record<number, number> = {}

    nums.forEach((num) => {
        const curr = countMap[num] ?? 0
        countMap[num] = curr + 1
    })

    const entries = Object.entries(countMap)

    for (let i = 0; i < entries.length; i++) {
        const [num, count] = entries[i]
        if (count === 1) {
            return Number(num)
        }
    }

    return 0
}