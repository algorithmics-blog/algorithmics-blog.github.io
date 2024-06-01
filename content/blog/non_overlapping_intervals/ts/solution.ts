export const eraseOverlapIntervals = (intervals: number[][]): number => {
    intervals.sort((a, b) => {
        return a[1] - b[1]
    })

    let res = 0
    let k = intervals[0][1]

    for (let i = 1; i < intervals.length; i++) {
        const end = intervals[i][1]
        const start = intervals[i][0]

        if (start >= k) {
            k = end
        } else {
            res++
        }
    }

    return res
}


