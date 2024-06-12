export const equalPairs = (grid: number[][]): number => {
    const countMap: Record<string, number> = {}
    let counter = 0

    grid.forEach(row => {
        let key = '';

        row.forEach(col => {
            key = `${key}${col};`
        })

        if (countMap[key]) {
            countMap[key]++
        } else {
            countMap[key] = 1
        }
    })

    for (let j = 0; j < grid.length; j++) {
        let key = '';

        for (let i = 0; i < grid.length; i++) {
            key = `${key}${grid[i][j]};`
        }

        if (countMap[key]) {
            counter += countMap[key]
        }
    }

    return counter
}
